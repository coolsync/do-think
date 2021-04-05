# day28



## advanced trait 1	关联 type	[Specifying Placeholder Types in Trait Definitions with Associated Types](https://doc.rust-lang.org/book/ch19-03-advanced-traits.html#specifying-placeholder-types-in-trait-definitions-with-associated-types)

```rust
// 1、关联类型在trait定义中指定占位符类型
// 关联类型是一个将类型占位符与trait相关联的方式。
// trait 的实现者会针对特定的实现在这个类型的位置指定相应的具体类型。
// 如此可以定义一个使用多种类型的 trait。
//pub trait Iterator {
//    type Item;
//    fn next(&mut self) -> Option<Self::Item>;
//}

pub trait Iterator1<T> {
    fn next(&mut self) -> Option<T>;
}

struct A{
    value: i32,
}

impl Iterator1<i32> for A {
    fn next(&mut self) -> Option<i32> {
        println!("in i32");
        if self.value > 3 {
            self.value += 1;
            Some(self.value)
        } else {
            None
        }
    }
}

impl Iterator1<String> for A {
    fn next(&mut self) -> Option<String> {
        println!("in String");
        if self.value > 3 {
            self.value += 1;
            Some(String::from("hello"))
        } else {
            None
        }
    }
}

fn main() {
    let mut a = A{value: 5};
    // a.next();   // error[E0282]: type annotations needed

    <A as Iterator1<i32>>::next(&mut a);    //完全限定语法，带上了具体的类型
    // println!("{}", a.value); // 6

    <A as Iterator1<String>>::next(&mut a);
    // println!("{}", a.value); // 7
    println!("Hello, world!");
}
```



## advanced trait 2	[Default Generic Type Parameters and Operator Overloading](https://doc.rust-lang.org/book/ch19-03-advanced-traits.html#default-generic-type-parameters-and-operator-overloading)

```rust
//2、默认泛型类型参数和运算符重载
//（1）使用泛型类型参数时，可以为泛型指定一个默认的具体类型。
//（2）运算符重载是指在特定情况下自定义运算符行为的操作。
//  Rust并不允许创建自定义运算符或者重载运算符，
//  不过对于std::ops中列出的运算符和相应的trait，我们可以实现运算符相关trait来重载。

use std::ops::Add;

#[derive(Debug, PartialEq)]
struct Point {
    x: i32,
    y: i32,
}

impl Add for Point {
    type Output = Point;
    fn add(self, other: Point) -> Point {
        Point {
            x: self.x + other.x,
            y: self.y + other.y,
        }
    }   
}

#[derive(Debug)]
struct MilliMeters(u32);
struct Meters(u32);

impl Add<Meters> for MilliMeters {
    type Output = MilliMeters;
    fn add(self, other:Meters) -> MilliMeters {
        MilliMeters(self.0 + other.0 * 1000)
    }
}

fn main() {
    assert_eq!(Point{x: 1, y: 1} + Point{x: 2, y: 2}, Point{x: 3, y: 3});
    
    let mi = MilliMeters(1);
    let m = Meters(1);
    let r = mi + m;
    println!("r: {:?}", r);
    println!("Hello, world!");
}

// pub trait Add<Rhs = Self> { //尖括号里面为默认类型参数，RHS是一个泛型类型参数（right hand side）
//     type Output;
//     fn add(self, rhs: Rhs) -> Self::Output;
// }
```



## advanced trait 3	[完全限定语法与消歧义：调用相同名称的方法](https://kaisery.github.io/trpl-zh-cn/ch19-03-advanced-traits.html#完全限定语法与消歧义调用相同名称的方法)



```rust
trait A {
    fn print(&self);
}

struct MyType;

impl A for MyType {
    fn print(&self) {
        println!("A trait for MyType")
    }
}

trait B {
    fn print(&self);
}

impl B for MyType {
    fn print(&self) {
        println!("B trait for MyType")
    }
}

impl MyType {
    fn print(&self) {
        println!("in MyType")
    }
}
fn main() {
    let my_type = MyType;
    my_type.print();    // eq: MyType::print(&my_type);
    A::print(&my_type);
    B::print(&my_type);

    println!("Hello, world!");
}
```



```rust
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
```





## advanced trait 4	[use super trait](https://doc.rust-lang.org/book/ch19-03-advanced-traits.html#using-supertraits-to-require-one-traits-functionality-within-another-trait)

```rust
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

impl OutPrint for Point {}  //  只到这一步,  err: `Point` doesn't implement `std::fmt::Display`

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
```







## advanced trait 5	 [use newtype pattern, implement external traits on external types](https://doc.rust-lang.org/book/ch19-03-advanced-traits.html#using-the-newtype-pattern-to-implement-external-traits-on-external-types)



```rust
// 5、newtype 模式用以在外部类型上实现外部 trait

// 孤儿规则（orphan rule）：只要 trait 或类型对于当前 crate 是本地的话就可以在此类型上实现该 trait。

// 一个绕开这个限制的方法是使用 newtype 模式（newtype pattern）。

use std::fmt;

// tuple struct
struct Wrapper(Vec<String>);

impl fmt::Display for Wrapper {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "({})", self.0.join(","))
    }
}
fn main() {
    let w = Wrapper(vec![String::from("hello"), String::from("world")]);
    println!("w : {}", w);    
}
```

