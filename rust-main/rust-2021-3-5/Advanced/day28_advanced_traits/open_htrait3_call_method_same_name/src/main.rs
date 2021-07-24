// 1 has self params
// trait A {
//     fn print(&self);
// }

// struct MyType;

// impl A for MyType {
//     fn print(&self) {
//         println!("A trait for MyType")
//     }
// }

// trait B {
//     fn print(&self);
// }

// impl B for MyType {
//     fn print(&self) {
//         println!("B trait for MyType")
//     }
// }

// impl MyType {
//     fn print(&self) {
//         println!("in MyType")
//     }
// }
// fn main() {
//     let my_type = MyType;
//     my_type.print();    // eq: MyType::print(&my_type);
//     A::print(&my_type);
//     B::print(&my_type);

//     println!("Hello, world!");
// }

// 2 not params
trait A {
    fn baby_name() -> String;
}

struct Dog;

impl Dog {
    fn baby_name() -> String {
        String::from("wangcai")
    }
}

impl A for Dog {
    fn baby_name() -> String {
        String::from("dahuang")
    }
}

fn main() {
    // let dog = Dog;
    // println!("baby name: {}", dog.baby_name()); // err: need self params

    println!("baby name: {}", Dog::baby_name());
    // println!("baby name: {}", A::baby_name());  // error
    println!("baby name: {}", <Dog as A>::baby_name()); //完全限定语法
}

//完全限定语法定义：
// <Type as Trait>::function(.....)