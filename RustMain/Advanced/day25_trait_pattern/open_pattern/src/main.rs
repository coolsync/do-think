// 1| 模式是Rust中特殊的语法，模式用来匹配值的结构。

// 2| A pattern consists of some combination of the following:

// Literals
// Destructured arrays, enums, structs, or tuples
// Variables
// Wildcards
// Placeholders

// match VALUE {
//     PATTERN => EXPRESSION,
//     PATTERN => EXPRESSION,
//     PATTERN => EXPRESSION,
// }

// pattern must match all codition
// fn main() {
//     let a = 1;
//     match a {
//         0 => println!("zero"),
//         1 => println!("one"),
//         _ => println!("other"),
//     }
//     println!("Hello, world!");
// }

// if let, has match, abort, back not run 
// fn main() {
//     let color: Option<&str> = None; // Mark

//     let is_ok = false;  
//     let age: Result<u8, _> = "33".parse();

//     if let Some(c) = color {
//         println!("color: {}", c);
//     } else if is_ok {
//         println!("is ok");
//     } else if let Ok(a) = age {
//         if a > 30 {
//             println!("mature man");
//         } else {
//             println!("young man");
//         }
//     } else {
//         println!("in else");
//     }
// }

// while let
//只要模式匹配就一直执行while循环
// fn main() {
//     let mut v = Vec::new();
//     v.push(1);
//     v.push(2);
//     v.push(3);

//     while let Some(top) = v.pop() {
//         println!("top: {}", top);
//     } // 只要匹配Some(value),就会一直循环
// }

// for
//在for循环中，模式是直接跟随for关键字的值，例如 for x in y，x就是对应的模式
// fn main() {
//     let v = vec!["a", "b", "c"];

//     for (index, value) in v.iter().enumerate() {    // enumerate get tuple
//         println!("index: {}, vlaue: {}", index, value);
//     }
// }
//此处的模式是(index, value)

// let
// let PATTERN = EXPRESSION
// fn main() {
// // (1, 2, 3) auto match (x, y, z), 1 bind to x, 2 bind to y,1 bind to z
//     let (x, y, z) = (1, 2, 3);
//     println!("{}, {}, {}", x, y, z);

//     let (x, .., z) = (1, 2, 3);
//     println!("{}, {}", x, z);
// }

// function
// function params 也是 模式
fn print_point(&(x, y): &(i32, i32)) {
    println!("x: {}, y: {}", x, y);
}

fn main() {
    let p = (1, 2);
    print_point(&p);
}

//模式在使用它的地方并不都是相同的，模式存在不可反驳的和可反驳的