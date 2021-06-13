// 1、泛型是具体类型或者其它属性的抽象替代，用于减少代码重复。
// 2、在Fn定义中使用泛型。
// 3、在结构体中使用泛型。
// 4、枚举中的泛型
// 5、方法中的泛型
// 例子2：方法和结构体中使用不同的类型
// 6、使用泛型并不会造成程序性能上的损失。rust通过在编译时进行泛型代码的单态化来保证效率。单态化时通过填充编译时使用的具体类型，将通用代码转换为特定代码的过程。

// // ----------No Generic-------------
// fn largest_i32(list: &[i32]) ->i32 {
//     let mut largest = list[0];
//     for &item in list.iter() {
//         if item > largest {
//             largest = item;
//         }
//     }
//     largest
// }

// fn largest_char(list: &[char]) -> char {
//     let mut largest = list[0];
//     for &item in list.iter() {
//         if item > largest {
//             largest = item;
//         }
//     }
//     largest
// }

// ---------- Generic ------------
// fn largest<T:PartialOrd + Copy>(list: &[T]) -> T {  //注意，要实现比较和复制的trait才行，否则报错
//     let mut larger = list[0];
//     for &item in list.iter() {
//         if item > larger {
//             larger = item;
//         }
//     }
//     larger
// }

// fn main() {
//     let num_list = vec![100,2,22,3,4,500];
//     // let num_max = largest_i32(&num_list);
//     let num_max = largest(&num_list);
//     println!("num max: {}", num_max);

//     let char_list = vec!['a', 'y', 'c', 'd'];
//     // let char_max = largest_char(&char_list);
//     let char_max = largest(&char_list);
//     println!("char max: {}", char_max);

//     println!("Hello, world!");
// }

// -------- In struct use generic --------
// #[derive(Debug)]
// struct Point<T> {
//     x: T,
//     y: T
// }

// #[derive(Debug)]
// struct Point2<T, U> {
//     x: T,
//     y: U
// }

// fn main(){
//     let p_int = Point{x: 1, y: 2};
//     println!("{:#?}", p_int);

//     let p_float = Point{x: 1.5, y: 2.5};
//     println!("{:?}", p_float);

//     let p = Point2{x: 1, y: 1.99};
//     println!("p: {:?}", p);
//     let p = Point2{x:1, y: 'c'};
//     println!("p: {:?}", p);
// }

// -------- Enum use generic --------
// enum Option<T> {
// 	Some(T),
// 	None,
// }

// enum Result<T, E> {
// 	Ok(T),
// 	Err(E),
// }

// -------- Method use generic --------
struct Point<T> {
    x: T,
    y: T,
}
impl<T> Point<T> {
    fn get_x(&self) -> &T {
        // no add &, move onwership
        &self.x
    }
    fn get_y(&self) -> &T {
        &self.y
    }
}

// 例子2：方法和结构体中使用不同的类型
#[derive(Debug)]
struct Point2<T, U> {
    x: T,
    y: U,
}

impl<T, U> Point2<T, U> {
    fn crate_point<V, W>(self, other:Point2<V, W>) -> Point2<T, W> {
        Point2 {
            x: self.x,
            y: other.y
        }
    }
}
fn main() {
    let p = Point { x: 1, y: 2 };
    println!("x: {}", p.get_x());
    println!("y: {}", p.get_y());

    let p = Point { x: 1.1, y: 2.2 };
    println!("x: {}", p.get_x());
    println!("y: {}", p.get_y());

    println!("++++++++++++++++++++++");

    let p1 = Point2{x: 1, y: 2.0};
    let p2 = Point2{x: "hello", y: 'a'};
    let p3 = p1.crate_point(p2);

    println!("p3: {:?}", p3)
}
