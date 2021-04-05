# day27



1、在此节之前讨论过的都是安全的Rust，即Rust在编译时会强制执行的内存安全保证。不会强制执行这类内存安全保证的，就是不安全的Rust。

2、不安全的Rust存在的两大原因：
（1）静态分析本质上是保守的，就意味着某些代码可能是合法的，但是Rust也会拒绝。在此情况下，可以使用不安全的代码。
（2）底层计算机硬件固有的不安全性。如果Rust不允许进行不安全的操作，有些任务根本就完成不了。

3、不安全的Rust具有的超级力量
Rust会通过unsafe关键字切换到不安全的Rust。不安全的Rust具有以下超级力量：
（1）解引用 raw 指针
（2）调用不安全的函数或者方法
（3）访问或修改可变静态变量
（4）实现不安全的trait
注意：unsafe并不会关闭借用检查器或禁用任何其它的Rust安全检查规则，它只提供上述几个不被编译器检查内存安全的功能。unsafe也不意味着块中的代码一定就是不ok的，它只是表示由程序员来确保安全。



## unsafe 1 [Dereferencing a Raw Pointer](https://doc.rust-lang.org/book/ch19-01-unsafe-rust.html#dereferencing-a-raw-pointer) 解引用 一个 原始指针



```rust
//4、解引用raw指针
//不可变和可变的，分别写作*const T, *mut T
//
//(1)允许忽略借用规则，可以同时拥有不可变和可变的指针，或者是多个指向相同位置的可变指针
//（2）不保证指向的内存是有效的
//（3）允许为空
//（4）不能实现任何自动清理的功能
//

fn main() {
    let mut num = 5;

    let r1 = &num as *const i32;
    let r2 = &mut num as *mut i32;

    unsafe {
        println!("r1: {}", *r1);
        println!("r2: {}", *r2);
    }

    // err: dereference of raw pointer is unsafe and requires unsafe function or block
    
    // println!("r1: {}", *r1);
    // println!("r2: {}", *r2);

    let add = 0x12456789usize;
    let _r = add as *const i32;

    println!("Hello, world!");
}
```





## unsafe 2 function

```rust
//5、调用 no safe 的 function 或者 method
unsafe fn dangerous() {
    println!("do something dangerous");
}

fn f1() {
    let mut num = 5;
    let r1 = &num as *const i32;
    let r2 = &mut num as *mut i32;

    unsafe {
        println!("*r1: {}", *r1);
        println!("*r2: {}", *r2);
    }
}
fn main() {
    unsafe {
        dangerous();
    }
    // dangerous(); // err

    f1();

    println!("Hello, world!");
}
```



## unsafe 3 Call C  Function

```rust
// call c lang function

extern "C" {
    fn abs(input: i32) -> i32;
}

fn main() {
    unsafe {
        println!("abs(-3): {}", abs(-3));
    }

    println!("Hello, world!");
}
```





## unsafe 4 Call rust Function in c

./by/src/lib.rc:

```rust
#![crate_type="staticlib"]	//  静态库
#[no_mangle]	// 不损坏
pub extern fn by() {
    println!("use rust");
}
```



by/Cargo.toml:

```toml
[package]
name = "by"
version = "0.1.0"
authors = ["coolsync <coolsyn@outlook.com>"]
edition = "2018"

# See more keys and their definitions at https://doc.rust-lang.org/cargo/reference/manifest.html

[dependencies]
[lib]
name = "by"
crate-type = ["staticlib"]
```



./main.c

```c
extern void by();	// 导入外部库
int main() {
    by();
    return 0;
}
```



./

```shell

$ chmod 755 libby.a 

$ gcc -o main main.c libby.a -lpthread -ldl

$ ./main
```





# day28



## unsafe 5  static variable



```rust
//6、访问或者修改可变静态变量
//
// static HELLO_WORLD: &str = "hello, world";

// fn main() {
//     println!("{}", HELLO_WORLD);
// }

//静态变量和常量的区别：
//1、静态变量有一个固定的内存地址（使用这个值总会访问相同的地址），常量则允许在任何被用到的时候复制其数据。
//2、静态变量可以是可变的，虽然这可能是不安全（用unsafe包含）

static mut COUNTER: u32 = 0;

fn add_counter(inc: u32) {
    unsafe {
        COUNTER += inc;
    }
}

fn main() {
    add_counter(3);
    add_counter(3);
    unsafe {
        println!("counter: {}", COUNTER);
    }
}
```







## unsafe 6 trait

```rust
//7、实现不安全的trait
//（1）当至少有一个方法中包含编译器不能验证的不变量时，该trait就是不安全的。
//（2）在trait之前增加unsafe声明其为不安全的，同事trait的实现也必须用unsafe标记。

unsafe trait A {
    fn a(&self);
}

struct B();

unsafe impl A for B {
    fn a(&self) {
        println!("hello");
    }
}

fn main() {
    let b = B();
    b.a();

    println!("Hello, world!");
}
```

















