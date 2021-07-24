//4、super trait 用于在另一个 trait 中使用某 trait 的功能

//有时我们可能会需要某个 trait 使用另一个 trait 的功能。
//在这种情况下，需要能够依赖相关的 trait 也被实现。
//这个所需的 trait 是我们实现的 trait 的 超 trait（supertrait）。


use std::fmt;

trait OutPrint: fmt::Display {
    fn out_print(&self) {
        let out_print = self.to_string();
        println!("out_print: {}", out_print)
    }
}

struct Point {
    x: i32,
    y: i32,
}

impl OutPrint for Point {}  //  只到这一步, err: `Point` doesn't implement `std::fmt::Display`

impl fmt::Display for Point {   // must impl supertrait fmt::Display
    //  fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "({}, {})", self.x, self.y)
    }
}

fn main() {
    let origin = Point { x: 0, y: 0 };
    assert_eq!(format!("The origin is: {}", origin), "The origin is: (0, 0)");

    println!("Hello, world!");
}
