# day30

## advanced type 1	[Type Aliases](https://doc.rust-lang.org/book/ch19-04-advanced-types.html#creating-type-synonyms-with-type-aliases)



```rust
//1、类型别名
type Kilometers = i32;

fn main() {
    let x: i32 = 5;
    let y: Kilometers = 6;
    let r: i32 = x + y;
    println!("x+y={}", r);
    println!("Hello, world!");
}

//类型别名的主要用途是减少重复。
// //（1）考虑如下类型：
// Box<dyn Fn() + Send + 'static>
// //如代码：
// let f: Box<dyn Fn() + Send + 'static> = Box::new(|| println!("hi"));

// fn takes_long_type(f: Box<dyn Fn() + Send + 'static>) {
//    // --snip--
// }

// fn returns_long_type() -> Box<dyn Fn() + Send + 'static> {
//    // --snip--
// }


// //使用别名，代码：
// type Thunk = Box<dyn Fn() + Send + 'static>;
// let f: Thunk = Box::new(|| println!("hi"));
// fn takes_long_type(f: Thunk) {
//    // --snip--
// }
// fn returns_long_type() -> Thunk {
//    // --snip--
// }

// //（2）考虑如下例子：
// use std::io::Error; //标准库中的std::io::Error结构体代表了所有可能的I/O错误
// use std::fmt;
// pub trait Write {
//    fn write(&mut self, buf: &[u8]) -> Result<usize, Error>;
//    fn flush(&mut self) -> Result<(), Error>;
//    fn write_all(&mut self, buf: &[u8]) -> Result<(), Error>;
//    fn write_fmt(&mut self, fmt: fmt::Arguments) -> Result<(), Error>;
// }

// //加上如下类型别名声明：
// type Result<T> = std::result::Result<T, std::io::Error>;//result<T, E> 中 E 放入了 std::io::Error

// //代码就可以变成：
// pub trait Write {
//    fn write(&mut self, buf: &[u8]) -> Result<usize>;
//    fn flush(&mut self) -> Result<()>;

//    fn write_all(&mut self, buf: &[u8]) -> Result<()>;
//    fn write_fmt(&mut self, fmt: Arguments) -> Result<()>;
// }
```







## advanced type 2	[The Never Type that Never Returns](https://doc.rust-lang.org/book/ch19-04-advanced-types.html#the-never-type-that-never-returns)



```rust
//2、从不返回的never type
//Rust 有一个叫做 ! 的特殊类型。在类型理论术语中，它被称为 empty type，因为它没有值。
//我们更倾向于称之为 never type。在函数不返回的时候充当返回值

use rand::Rng;
use std::cmp::Ordering;
use std::io;

// 1. 生成随机数
// 2. 多次提示 user input
// 3. 创建 guess String 字符串，
// 保存用户输入内容
// 4. 转换 guess 成 u32, has err, 跳出此次 loop
// 5. 使用 cmp包 比较 guess 与 secret_num 是否一致
// 6, win, 使用 break 跳出 loop
fn main() {
    println!("Guess a num Game");

    let secret_num = rand::thread_rng().gen_range(1..101);

    println!("secret_num: {}", secret_num);

    loop {
        println!("Please a num: ");

        let mut guess = String::new();

        io::stdin()
            .read_line(&mut guess)
            .expect("read input content faild");

        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => continue, //continue 的值是 !。
                                //当 Rust 要计算 guess 的类型时，它查看这两个分支。
                                //前者是 u32 值，而后者是 ! 值。
                                //因为 ! 并没有一个值，Rust 决定 guess 的类型是 u32
        };

        match guess.cmp(&secret_num) {
            Ordering::Less => println!("too less"),
            Ordering::Greater => println!("too greater"),
            Ordering::Equal => {
                println!("you win");
                break;
            }
        }
    }
}


// // 例子2：panic!
// // Option 上的 unwrap 函数代码：

// impl<T> Option<T> { 
// 	pub fn unwrap(self) -> T { 
// 		match self { 
// 			Some(val) => val, 
// 			None => panic!("called `Option::unwrap()` on a `None` value"),
// 		} 
// 	}
// }

// // 说明：
// // match 时，Rust 知道 val 是 T 类型，panic! 是 ! 类型，所以整个 match 表达式的结果是 T 类型。
```



## advanced type 3	[Dynamically Sized Types and the `Sized` Trait](https://doc.rust-lang.org/book/ch19-04-advanced-types.html#dynamically-sized-types-and-the-sized-trait)



```rust
// 3、动态大小类型和Sized trait
// 动态大小类型（dynamically sized types），有时被称为 “DST” 或 “unsized types”，
// 这些类型允许我们处理只有在运行时才知道大小的类型。
// (1)最典型的就是str

fn main() {
    let s1: &str = "hello";
    let s2: &str = "hello";

    // let s1: str = "hello";  // the size for values of type `str` cannot be known at compilation time
    // let s2: str = "hello";
    println!("Hello, world!");
}

//&str有两个值：str的地址和长度, &str的大小在编译时就可以知道，长度就是2*usize
//动态大小类型的黄金规则：必须将动态大小类型的值置于某种指针之后，如：Box<str>\Rc<str>

//另一个动态大小类型是trait。每一个 trait 都是一个可以通过 trait 名称来引用的动态大小类型。
//为了将 trait 用于 trait 对象，必须将他们放入指针之后，
//比如 &Trait 或 Box<Trait>（Rc<Trait> 也可以）。




////（2） Sized trait
////为了处理 DST，Rust 用Sized trait 来决定一个类型的大小是否在编译时可知。
////这个 trait 自动为编译器在编译时就知道大小的类型实现。
////例子：
//
//fn generic<T>(t: T) {//T为编译时就知道大小的类型
//    // --snip--
//}
//
////等价于
//fn generic<T: Sized>(t: T) {//T为编译时就知道大小的类型
//    // --snip--
//}
//
////如何放宽这个限制呢？Rust提供如下方式：
//fn generic<T: ?Sized>(t: &T) {//T 可能是Sized，也可能不是 Sized 的
//    // --snip--
//}

```

