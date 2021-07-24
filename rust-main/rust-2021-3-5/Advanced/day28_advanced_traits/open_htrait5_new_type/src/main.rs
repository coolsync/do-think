// 5、newtype 模式用以在外部类型上实现外部 trait

// 孤儿规则（orphan rule）： 它说明只要 trait 或类型对于当前 crate 是本地的话就可以在此类型上实现该 trait。

// 一个绕开这个限制的方法是使用 newtype 模式（newtype pattern）。

use std::fmt;

// Tuple struct
struct Wrapper(Vec<String>);    // Vec<String> is out type

impl fmt::Display for Wrapper { // fmt::Display is out trait
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "({})", self.0.join(","))
    }
}
fn main() {
    let w = Wrapper(vec![String::from("hello"), String::from("world")]);
    println!("w : {}", w);    
}
