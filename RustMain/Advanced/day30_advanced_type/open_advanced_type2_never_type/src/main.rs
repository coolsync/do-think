//2、从不返回的never type
//Rust 有一个叫做 ! 的特殊类型。在类型理论术语中，它被称为 empty type，因为它没有值。
//我们更倾向于称之为 never type。在函数不返回的时候充当返回值

use std::io;

fn main() {
    println!("Guess the number!");
    println!("Hello, world!");
}
