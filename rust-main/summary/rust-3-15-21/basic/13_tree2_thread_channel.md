# day22



## week tree 2



### 强引用 与 弱引用 区别：

```rust
use std::rc::{Rc, Weak};
use std::cell::RefCell;

struct Node {
    value: i32,
    parent: RefCell<Weak<Node>>,
    children: RefCell<Vec<Rc<Node>>>,
}

fn main() {
    let leaf = Rc::new(Node{
        value: 3,
        parent: RefCell::new(Weak::new()),
        children: RefCell::new(vec![]),
    });
    println!(
        "1 leaf strong: {}, weak: {}",
        Rc::strong_count(&leaf), Rc::weak_count(&leaf)
    );
    {
        let branch = Rc::new(Node{
            value: 5,
            parent: RefCell::new(Weak::new()),
            children: RefCell::new(vec![Rc::clone(&leaf)]), // branch -> leaf 
        });

        println!(
            "1 branch strong: {}, weak: {}",
            Rc::strong_count(&branch), Rc::weak_count(&branch)
        );

        *leaf.parent.borrow_mut() = Rc::downgrade(&branch); // leaf <- branch

        println!(
            "2 leaf strong: {}, weak: {}",
            Rc::strong_count(&leaf), Rc::weak_count(&leaf)
        );
        println!(
            "2 branch strong: {}, weak: {}",
            Rc::strong_count(&branch), Rc::weak_count(&branch)
        );
    }
    
    println!(
        "3 leaf strong: {}, weak: {}",
        Rc::strong_count(&leaf), Rc::weak_count(&leaf)
    );
    
    println!("Hello, world!"); 
}
```



result:

```bash
1 leaf strong: 1, weak: 0
1 branch strong: 1, weak: 0

2 leaf strong: 2, weak: 0
2 branch strong: 1, weak: 1

3 leaf strong: 1, weak: 0

Hello, world!
```



## thread 1



（1）进程是资源分配的最小单位，线程是CPU调度的最小单位。

（2）在使用多线程时，经常会遇到的一些问题：

​         1. 竞争状态：多个线程以不一致的顺序访问数据或资源；

​          2.死锁：两个线程相互等待对方停止使用其所拥有的资源，造成两者都永久等待；A: 1->2->3  B: 2->1->3    

t1: A:1, B:2    接下来： A：2， B：1   造成死锁

​          3.只会发生在特定情况下且难以稳定重现和修复的bug

（3）编程语言提供的线程叫做绿色线程，如go语言，在底层实现了M:N的模型，即M个绿色线程对应N个OS线程。但是，Rust标准库只提供1：1的线程模型的实现，即一个Rust线程对应一个Os线程。

​         运行时代表二进制文件中包含的由语言本身提供的代码，这些代码根据语言的不同可大可小，不过非汇编语言都会有一定数量的运行时代码。通常，大家说一个语言“没有运行时”，是指这个语言的“运行时”很小。Rust、C都是几乎没有运行时的。



```rust
use std::thread;
use std::time::Duration;

fn main() {
    // create a thread
    let handle = thread::spawn(|| {
        for i in 0..9 {
            println!("num {}, spawn thread", i);
            thread::sleep(Duration::from_millis(1));
        }
    });
    
    // Wait spawn thread end
    handle.join().unwrap();
    
    for i in 0..5 {
        println!("num {}, main thread", i);
        thread::sleep(Duration::from_millis(1));
    }
    // wait spawn thread end
    // handle.join().unwrap();
}
```



## thread 2 move closures



线程 与 move 闭包

```rust
fn main() {
    let v = vec![1, 2, 3];

    let handle = thread::spawn(|| {
        println!("v {:?}", v);
    });
    
    // drop(v)

    handle.join().unwrap();
    
    println!("Hello, world!");
}
```

res:

```bash
error[E0382]: borrow of moved value: `v`
  --> src/main.rs:12:24
   |
4  |     let v = vec![1, 2, 3];
   |         - move occurs because `v` has type `Vec<i32>`, which does not implement the `Copy` trait
5  | 
6  |     let handle = thread::spawn( move || {
   |                                 ------- value moved into closure here
7  |         println!("v {:?}", v);
   |                            - variable moved due to use in closure
...
12 |     println!("v {:?}", v);
   |                        ^ value borrowed here after move

error: aborting due to previous error

```



use move:

```rust
use std::thread;

fn main() {
    let v = vec![1, 2, 3];

    let handle = thread::spawn( move || {
        println!("v {:?}", v);	// v move here
    });

    handle.join().unwrap();
    
    // println!("v {:?}", v); // not print
    
    println!("Hello, world!");
}
```



# day23



## Channel 1

```rust
use std::thread;
use std::sync::mpsc;
use std::time::Duration;

fn main() {
    let (tx, rx) = mpsc::channel();
    
    thread::spawn(move || {
        let val = String::from("hi");
        thread::sleep(Duration::from_secs(3));  // after 3s send
        tx.send(val).unwrap();
    });

    let recieved = rx.recv().unwrap();
    println!("Recv Get: {}", recieved);
}


// Konw：
//1、发送者的send方法返回的是一个Result<T,E>,
//如果接收端已经被丢弃了，将没有发送值的目标，此时发送会返回错误。
//2、接受者的recv返回值也是一个Result类型，当通道发送端关闭时，返回一个错误值。
//3、接收端这里使用的recv方法，会阻塞到有一个消息到来。我们也可以使用try_recv()，不会阻塞，会立即返回。

//1、Rust中一个实现消息传递并发的主要工具是通道。通道由两部分组成，一个是发送端，一个是接收端，发送端用来发送消息，接收端用来接收消息。发送者或者接收者任一被丢弃时就可以认为通道被关闭了。
//
//2、通道介绍
//（1）通过mpsc::channel，创建通道，mpsc是多个生产者，单个消费者；
//（2）通过spmc::channel，创建通道，spmc是一个生产者，多个消费者；
//（3）创建通道后返回的是发送者和消费者，示例：
//let (tx, rx) = mpsc::channel();
//let (tx, rx) = spmc::channel();
```



## Channel 2



```rust
use std::thread;
use std::sync::mpsc;

fn main() {
    // create channel
    let (tx, rx) = mpsc::channel();

    thread::spawn(move || {
        let val = String::from("hi");
        tx.send(val).unwrap(); 
        println!("val: {}", val); //调用send的时候，会发生move动作，所以此处不能再使用val
    });

    let re = rx.recv().unwrap();
    println!("Recv Get: {}", re);
}

```



bug:

```bash
error[E0382]: borrow of moved value: `val`
  --> src/main.rs:11:29
   |
9  |         let val = String::from("hi");
   |             --- move occurs because `val` has type `String`, which does not implement the `Copy` trait
10 |         tx.send(val).unwrap(); 
   |                 --- value moved here
11 |         println!("val: {}", val);
   |                             ^^^ value borrowed here after move
```



## channel 3



```rust
use std::sync::mpsc;
use std::thread;
use std::time::Duration;

fn main() {
    let (tx, rx) = mpsc::channel();

    thread::spawn(move || {
        let vals = vec![
            String::from("hi"),
            String::from("from"),
            String::from("the"),
            String::from("thread"),
        ];

        for val in vals {
            tx.send(val).unwrap();
            thread::sleep(Duration::from_secs(1));
        }
    });

    // for recv in rx.recv() {  // rx.recv() only recv a val
    for recv in rx {
        println!("Get: {}", recv);
    }

    println!("Hello, world!");
}
```



## channel 4

```rust
use std::thread;
use std::sync::mpsc;
use std::time::Duration;

fn main() {
    // mutiple produce single custom
    let (tx, rx) = mpsc::channel();
    let tx1 = mpsc::Sender::clone(&tx);
    let tx2 = mpsc::Sender::clone(&tx);

    thread::spawn(move || {
        let vals = vec![
            String::from("hi"),
            String::from("from"),
            String::from("the"),
            String::from("thread"),
        ];

        for val in vals {
            tx.send(val).unwrap();  // not use tx, main thread still wait
            thread::sleep(Duration::from_secs(1));
        }
    });

    thread::spawn(move || {
        let vals = vec![
            String::from("A"),
            String::from("B"),
            String::from("C"),
            String::from("D"),
        ];

        for val in vals {
            tx1.send(val).unwrap();
            thread::sleep(Duration::from_secs(1));
        }
    });

    thread::spawn(move || {
        let vals = vec![
            String::from("aa"),
            String::from("bb"),
            String::from("cc"),
            String::from("dd"),

        ];
        
        for val in vals {
            tx2.send(val).unwrap();
            // thread::sleep(Duration::from_secs(1));
        }
    });

    for recv in rx {
        println!("Recv Get: {}", recv);
    }

    println!("Hello, world!");
}
```



result:

```bash
Recv Get: hi
Recv Get: A
Recv Get: aa
Recv Get: bb
Recv Get: cc
Recv Get: dd
Recv Get: from
Recv Get: B
Recv Get: the
Recv Get: C
Recv Get: thread
Recv Get: D
Hello, world!
```

