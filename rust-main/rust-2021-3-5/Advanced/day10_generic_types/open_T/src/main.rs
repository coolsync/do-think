// 1| Generic types: 是具体类型 or 其他属性的 抽象替代， 用于减少代码的重复
// 2| Function: Removing Duplication by Extracting a Function
// 3| Struct: In Struct Definitions, use generics

// // ------------function--------------
// // ---------no generic types---------
// fn largest_i32(list: &[i32]) -> i32 {
//     let mut largest = list[0];

//     for &item in list {
//         if item > largest {
//             largest = item;
//         }
//     }

//     largest
// }

// fn largest_char(list: &[char]) -> char {
//     let mut largest = list[0];

//     for &item in list {
//         if item > largest {
//             largest = item;
//         }
//     }

//     largest
// }

// // ---------generic types---------
// fn largest<T: PartialOrd + Copy>(list: &[T]) -> T {
//     let mut larger = list[0];

//     for &item in list {
//         if item > larger {
//             larger = item
//         }
//     }

//     larger
// }

// fn main() {
//     let number_list = vec![34, 50, 25, 100, 65];

//     // let result = largest_i32(&number_list);
//     let result = largest(&number_list);
//     println!("largest number: {}", result);

//     let char_list = vec!['y', 'm', 'a', 'q'];

//     // let result = largest_char(&char_list);
//     let result = largest(&char_list);
//     println!("largest char: {}", result);
// }

// // ------------struct generic--------------
// #[derive(Debug)]
// struct Point<T> {
//     x: T,
//     y: T,
// }

// #[derive(Debug)]
// struct Point2<T, U> {
//     x: T,
//     y: U,
// }

// fn main() {
//     let integer = Point { x: 1, y: 2 };
//     println!("{:#?}", integer);

//     let float = Point { x: 1.0, y: 5.0 };
//     println!("{:?}", float);

//     let a = Point2 { x: 1.0, y: 'a' };
//     println!("{:?}", a);

//     let both_integer = Point2 { x: 5, y: 10 };
//     let both_float = Point2 { x: 3.1, y: 7.1 };
//     let integer_and_float = Point2 { x: 5, y: 4.0 };
//     println!("{:?}", integer_and_float);
// }

// ------------enum generic--------------
// enum Option<T> {
// Some(T),
// None,
// }
//
// enum Result<T, E> {
// Ok(T),
// Err(E),
// }

// ------------struct method generic--------------
struct Point<T> {
    x: T,
    y: T,
}

impl<T> Point<T> {
    fn get_x(&self) -> &T {
        &self.x
    }

    fn get_y(&self) -> &T {
        &self.y
    }
}

struct Point2<T, U> {
    x: T,
    y: U,
}

impl<T, U> Point2<T, U> {
    fn crate_point<V, W>(self, other: Point2<V, W>) -> Point2<T,W> {
        Point2 {
            x: self.x,
            y: other.y,
        }
    }
}

fn main() {
    // let p1 = Point { x: 1, y: 2 };
    // println!("p1.x = {}, p2.y = {}", p1.get_x(), p1.get_y());

    let p1 = Point2 { x: 5, y: 10.2 };
    let p2 = Point2 { x: "hello", y: 'c' };

    let p3 = p1.crate_point(p2);

    println!("p3.x = {}, p3.y = {}", p3.x, p3.y)
}
