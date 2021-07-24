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

// Change Add default type Rhs
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

