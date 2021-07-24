# Guessing Number



println!

let

mut	// mutable	// immutable

use std::io

io.stdin()

read_line()

expect()

{}



```rust
use std::io;

fn main() {
    println!("Guess the number!");

    println!("Please input your guess.");

    let mut guess = String::new();

    io::stdin()
        .read_line(&mut guess)
        .expect("Failed to read line");

    println!("You guessed: {}", guess);
}
```



## [Generating a Secret Number](https://doc.rust-lang.org/book/ch02-00-guessing-game-tutorial.html#generating-a-secret-number)



Cargo.toml

```toml
[dependencies]
rand = "0.8.0"
```

cargo build

If the `rand` crate has released two new versions, `0.5.6` and `0.6.0`, you would see the following if you ran `cargo update`:

```console
$ cargo update
    Updating crates.io index
    Updating rand v0.5.5 -> v0.5.6
```



std::cmp::Ordering

let guess: u32

guess.parse().expect("...")

match

guess.cmp(&number)

Ordering::Less

Ordering::Greater

Ordering::Equal

```rust
use std::io;
use rand::Rng;  // trit
use std::cmp::Ordering;

fn main() {  
    // gen scret number
    let scret_number = rand::thread_rng().gen_range(1..101);
    
    println!("scret number:{}", scret_number);
    
    println!("Guessing a number.");

    let mut guess = String::new();
    
    io::stdin()
    .read_line(&mut guess)
    .expect("read line err");
    
    println!("you guess num: {}", guess);
    
    // trim 消除边缘多余的stuff
    let guess: u32 = guess.trim().parse().expect("you sld input a num");
    
    match guess.cmp(&scret_number) {
        Ordering::Less => println!("too samll."),
        Ordering::Greater => println!("too big"),
        Ordering::Equal => println!("you win")
    }
}
```



## [Handling Invalid Input](https://doc.rust-lang.org/book/ch02-00-guessing-game-tutorial.html#handling-invalid-input)



```rust
    // gen scret number
    let scret_number = rand::thread_rng().gen_range(1..101);
    
    println!("scret number:{}", scret_number);

    loop {
        println!("Guessing a number.");
    
        let mut guess = String::new();
        
        io::stdin().read_line(&mut guess).expect("read line err");
        
        println!("you guess num: {}", guess);
        
        // trim 消除边缘多余的stuff
        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            // Err(_) => continue, // err no handle
            Err(_) => {
                println!("parse err");
                break
            },
        };
        
        match guess.cmp(&scret_number) {
            Ordering::Less => println!("too samll."),
            Ordering::Greater => println!("too big"),
            Ordering::Equal => {
                println!("you win");
                break
            }
        }
    }
```





# [Common Programming Concepts](https://doc.rust-lang.org/book/ch03-00-common-programming-concepts.html#common-programming-concepts)



## [Variables and Mutability](https://doc.rust-lang.org/book/ch03-01-variables-and-mutability.html#variables-and-mutability)



```rust
  let mut x = 5;
    println!("The value of x is: {}", x);

    // x = 6;  // if `let x = 5;` cannot assign twice to immutable variable

    x = 6;
    println!("The value of x is: {}", x);
```



### [Differences Between Variables and Constants](https://doc.rust-lang.org/book/ch03-01-variables-and-mutability.html#differences-between-variables-and-constants)



```rust
let mut spaces = "    ";    // variable does not need to be mutable
    let spaces = spaces.len();

    
    println!("The value of spaces is: {}", spaces)
```



```rust
fn main() {
    another_function(5, 6); //arguments   
}

fn another_function(x: i32, y: i32){   // parameters
    println!("the value of x is: {}", x);
    println!("the value of y is: {}", y);
}
```



# [Understanding Ownership](https://doc.rust-lang.org/book/ch04-00-understanding-ownership.html#understanding-ownership)

1  rust 通过 ownership 机制来管理内存，编译器在编译时由 ownership 机制对内存使用进行check

2  heap and stack

  compile time data type space  fixed, allocation to stack

  compile time, data type space not fixed, allocation to heap

3  scope 作用域

4  String 内存回收

5  move

6  clone

7  stack data copy

8  function and scope



heap

```rust
{
        let mut s1 = String::from("hello");
        s1.push_str(", world");	// heap
        println!("s1: {}", s1);
}
```

move

```rust
{
        // let mut s1 = String::from("hello");
        // s1.push_str(", world");

        let s1 = String::from("hello");
        println!("s1: {}", s1);

        let s2 = s1;
        println!("s2: {}", s2); // borrow of moved value: `s1` 
        // println!("s1: {}", s1); // move to s2, s1 invalid, value borrowed here after move
}
```



clone

```rust
{
        // let mut s1 = String::from("hello");
        // s1.push_str(", world");

        let s1 = String::from("hello");
        println!("s1: {}", s1);

        let s2 = s1;
        println!("s2: {}", s2); // borrow of moved value: `s1` 
        // println!("s1: {}", s1); // move to s2, s1 invalid, value borrowed here after move
    
        let s3 = s2.clone();
        println!("s3: {}", s3);
        println!("s2: {}", s2);
}
```



 stack data copy

```rust
 	// copy trait
    let a = 1;
    let b = a;
    println!("a: {}, b: {}", a, b);
    //常用的具有copy trait有：
    //所有的整型
    //浮点型
    //布尔值
    //字符类型 char
    //元组
```



function and scope

```rust
fn takes_ownership(something: String) -> String {
    println!("{}", something); 
    something   // scope in main fn
} 

fn make_cp(i :i32) {
    println!("{}", i);
}
fn main() {
    let s = String::from("nice");   // heap
    // takes_ownership(s); // invalid, scope in takes_ownership fn
    // println!("{}", s); 

    let s1 = takes_ownership(s);
    println!("{}", s1);

    let i = 1;  // stack
    make_cp(i);
    println!("i = {}", i);  // valid
}
```





## [References and Borrowing](https://doc.rust-lang.org/book/ch04-02-references-and-borrowing.html#references-and-borrowing)



```rust
fn main() {
    let s1 = gives_onwership();
    println!("s1 = {}", s1);  

    let s2 = String::from("hello s2");

    let s3 = takes_and_gives_back(s2);

    // println!("s2 = {}", s2); // err, this value move, again ref it
}

fn gives_onwership() -> String {
    let s = String::from("hello s1");
    s  
}

fn takes_and_give_back(s: String) -> String {
    s
}
```



```rust
fn main() {
    let s1 = gives_onwership();
    println!("s1 = {}", s1);  

    let mut s2 = String::from("hello s2");

    let s3 = takes_and_give_back(s2);

    s2 = takes_and_give_back(s3);
    
    println!("s2 = {}", s2);    //  let mut s2, ok

    // println!("s2 = {}", s2); // err, let s2, this value move, again ref it
}

fn gives_onwership() -> String {
    let s = String::from("hello s1");
    s  
}

fn takes_and_give_back(s: String) -> String {
    s
}
```



### ref and borrow

ref 引用
引用: 用法 `&`,
让我们创建一个指向值的引用，但是并不拥有它，因为不拥有这个值，所以，当引用离开其值指向的作用域后也不会被丢弃

```rust


fn calcute_length(s: &String) -> usize {
    s.len()
}

fn main() {
    let s1 = String::from("hello");    
    
    let s = &s1;

    let len = calcute_length(s);

    println!("s1 = {}", s1);  // can use

    println!("len = {}", len);
}
```



借用: `&mut`

```rust

fn calcute_length(s: &String) -> usize {
    s.len()
}

fn modify_string(s: &mut String) {
    s.push_str(", world");
}

fn main() {
    let mut s1 = String::from("hello");    
    
    // let s = &s1;
    // let len = calcute_length(s);
    // println!("s = {}", s);  // can use
    // println!("len = {}", len);

    modify_string(&mut s1);

    println!("s1 = {}", s1);      
}
```







```rust
let mut s1 = String::from("hello");    
    
    // let s = &s1;
    // let len = calcute_length(s);
    // println!("len = {}", len);

    // modify_string(&mut s1);

    let r1 = &s1;
    let r2 = &s1;
    let r3 = &mut s1;

    println!("{}, {}, {}", r1, r2, r3);
    // println!("s1 = {}", s1);     
```



```bash
error[E0502]: cannot borrow `s1` as mutable because it is also borrowed as immutable
  --> main.rs:50:14
   |
48 |     let r1 = &s1;
   |              --- immutable borrow occurs here
49 |     let r2 = &s1;
50 |     let r3 = &mut s1;
   |              ^^^^^^^ mutable borrow occurs here
51 | 
52 |     println!("{}, {}, {}", r1, r2, r3);
   |                            -- immutable borrow later used here
```



solution:

```rust
let mut s1 = String::from("hello");

let r1 = &s1;
let r2 = &s1;
println!("{},{}", r1, r2);
    
let r3 = &mut s1;
println!("{}", r3);
   // println!("s1 = {}", s1);
```







# slice



```rust
//1、字符串slice是String中一部分值的引用
//2、字面值就是slice
//3、其它类型slice

fn main() {
    let s = String::from("hello world");

    let h = &s[0..5];
    let h = &s[0..=4];
    let h = &s[..=4];
    let h = &s[..5];
    println!("h = {}", h);

    let w = &s[6..11];
    let w = &s[6..];
    let w = &s[6..=10];
    let w = &s[..];
    println!("w = {}", w);

    let ss = String::from("你好");

    // let nihao = &ss[0..1];
    // thread 'main' panicked at 'byte index 1 is not a char boundary; it is inside '你' (bytes 0..3) of `你好`'
    // println!("nihao = {}", nihao);

    // let s1 = "hhh";    // &str
    // let h1 = s1[0..1];
    // println!("h1 = {}", h1);  // the type `str` cannot be indexed by `{integer}` string indices are ranges of `usize`

    let a = [1, 2, 3, 4, 5];

    let aa = &a[1..3];  // 2, 3
    println!("aa = {}", aa[0]);
    println!("aa = {}", aa[1]);
}
```





