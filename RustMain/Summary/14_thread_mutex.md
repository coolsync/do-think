# day24



## thread 3 mutex



```rust
//1、通道类似于单所有权的方式，值传递到通道后，发送者就无法再使用这个值；
//2、共享内存类似于多所有权，即多个线程可以同时访问相同的内存位置。

//互斥器：mutex
//1、任意时刻，只允许一个线程来访问某些数据;
//2、互斥器使用时，需要先获取到锁，使用后需要释放锁。
// Mutux<T>

use std::sync::Mutex;

fn main() {
    let m = Mutex::new(5);
    {
        let mut num = m.lock().unwrap();
        *num = 6;
    }   // leave scope, auto release // 离开作用域时，自动释放

    println!("m: {:?}", m);
}

//Mutex<T>是一个智能指针，lock调用返回一个叫做MutexGuard的智能指针
//内部提供了drop方法，实现当MutexGuard离开作用域时自动释放锁。
```



## thread 4 mutex counter



### 1 counter and mutex

```rust
use std::thread;
use std::sync::Mutex;

fn main() {
    let counter = Mutex::new(0);
    let handles = vec![];

    for _ in 0..10 {
        let handle = thread::spawn(move || {
            let mut num = counter.lock().unwrap();
            *num += 1;
        });

        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap(); // wait 对应 thread end
    }

    println!("result: {}", *counter.lock().unwrap());
}
```



bug:

```bash
error[E0382]: use of moved value: `counter`
  --> src/main.rs:10:36
   |
6  |     let counter = Mutex::new(0);
   |         ------- move occurs because `counter` has type `Mutex<i32>`, which does not implement the `Copy` trait
...
10 |         let handle = thread::spawn(move || {
   |                                    ^^^^^^^ value moved into closure here, in previous iteration of loop
11 |             let mut num = counter.lock().unwrap();
   |                           ------- use occurs due to use in closure

error: aborting due to previous error; 
```





### 2 use `rc::Rc` 实现 多个 可读 共享  Err

```rust
use std::sync::Mutex;
use std::thread;
use std::rc::Rc;

fn main() {
    // let counter = Mutex::new(0);
    let counter = Rc::new(Mutex::new(0));  // `Rc<Mutex<i32>>` cannot be sent between threads safely
    let mut handles = vec![];

    for _ in 0..10 {
        let cnt = Rc::clone(&counter);
        let handle = thread::spawn(move || {
            let mut cnt = counter.lock().unwrap();
            *cnt += 1;
        });

        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap(); // wait 对应 thread end
    }

    println!("result: {}", *counter.lock().unwrap());
}
```





bug:

```shell
error[E0277]: `Rc<Mutex<i32>>` cannot be sent between threads safely
   --> src/main.rs:12:22
    |
12  |           let handle = thread::spawn(move || {
    |  ______________________^^^^^^^^^^^^^_-
    | |                      |
    | |                      `Rc<Mutex<i32>>` cannot be sent between threads safely
13  | |             let mut cnt = counter.lock().unwrap();
14  | |             *cnt += 1;
15  | |         });
    | |_________- within this `[closure@src/main.rs:12:36: 15:10]`
    | 
   ::: /home/dart/.rustup/toolchains/stable-x86_64-unknown-linux-gnu/lib/rustlib/src/rust/library/std/src/thread/mod.rs:617:8
    |
617 |       F: Send + 'static,
    |          ---- required by this bound in `spawn`
    |
    = help: within `[closure@src/main.rs:12:36: 15:10]`, the trait `Send` is not implemented for `Rc<Mutex<i32>>`
    = note: required because it appears within the type `[closure@src/main.rs:12:36: 15:10]`
```



### 3 use `sync::Arc` 实现 多个 可读 共享 

```rust
//RefCell\Rc\Box

//RefCell<T>/Rc<T> 与Mutex<T>/Arc<T>
//1、Mutex<T>提供内部可变性，类似于RefCell
//2、RefCell<T>/Rc<T>是非线程安全的， Mutex<T>/Arc<T>是线程安全的

// use std::rc::Rc;
use std::sync::Mutex;
use std::thread;
use std::sync::Arc;

fn main() {
    // let counter = Mutex::new(0);
    // let counter = Rc::new(Mutex::new(0)); // `Rc<Mutex<i32>>` threads is not safely
    
    let counter = Arc::new(Mutex::new(0));
    let mut handles = vec![];

    for _ in 0..10 {
        let ct = Arc::clone(&counter);
        let handle = thread::spawn(move ||{
           let mut num = ct.lock().unwrap();
           *num += 1; 
        });

        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();     // wait 对应 thread end
    }

    println!("result: {:?}", counter);
}
```



rs:

```shell
Finished dev [unoptimized + debuginfo] target(s) in 0.01s
Running `target/debug/open_thread4_mutex`
result: Mutex { data: 10 }
```







1、有两个并发概念内嵌于语言中：std::marker中的Sync和Send trait。

2、通过Send允许在线程间转移所有权
（1）Send标记trait表明类型的所有权可以在线程间传递。几乎所有的Rust类型都是Send的，但是例外：例如Rc<T>是不能Send的。
（2）任何完全由Send类型组成的类型也会自动被标记为Send。
//struct A {
//	a
//	b
//	c
//}

3、Sync允许多线程访问
（1）Sync 标记 trait 表明一个实现了 Sync 的类型可以安全的在多个线程中拥有其值的引用，即，对于任意类型 T，如果 &T（T 的引用）是 Send 的话 T 就是 Sync 的，这意味着其引用就可以安全的发送到另一个线程。
（2）智能指针 Rc<T> 也不是 Sync 的，出于其不是 Send 相同的原因。RefCell<T>和 Cell<T> 系列类型不是 Sync 的。RefCell<T> 在运行时所进行的借用检查也不是线程安全的，Mutex<T> 是 Sync 的。

4、手动实现Send和Sync是不安全的
通常并不需要手动实现 Send 和 Sync trait，因为由 Send 和 Sync 的类型组成的类型，自动就是 Send 和 Sync 的。因为他们是标记 trait，甚至都不需要实现任何方法。他们只是用来加强并发相关的不可变性的。`