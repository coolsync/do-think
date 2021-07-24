// Costum Box 
use std::ops::Deref;

struct MyBox<T>(T);

impl<T> MyBox<T> {
    fn new(x: T) -> MyBox<T> {
        MyBox(x)
    }
}

impl<T> Deref for MyBox<T> {
    type Target = T;
    fn deref(&self) -> &T {
        &self.0
    }
}

fn main() {
    let x = 5;
    let y = MyBox::new(x);

    assert_eq!(5, x);
    assert_eq!(5, *y);

    let m = MyBox::new(String::from("Rust"));
    hello(&m); // 将 MyBox 变为 &String，再将 String 解引用，变为字符串 slice。  &str
}

fn hello(name: &str) {
    println!("Hello, {}", name);
}

// 解引用多态与可变性交互
// when T: Defer<Target:U>, from &T to &U,
// when T: Defer<Target:U>, from &mut T to &U, 
// when T: DeferMut<Target:U>, from &mut T to &mut U,



