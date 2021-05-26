// 6、忽略模式中的值

// fn aa(_: i32, y: i32) {
//     println!("y: {}", y)
// }

// trait A {
//     fn bb(x: i32, y: i32);
// }

// struct B {}

// impl A for B {
//     fn bb(_: i32, y: i32) {
//         println!("y: {}", y)
//     }
// }

// fn main() {
//     aa(1, 2);
//     let numbers = (1, 2, 3, 4);

//     match numbers {
//         (one, _, three, _) => {
//             println!("one: {}, three: {}", one, three);
//         }
//     }
//     println!("Hello, world!");
// }

fn main() {
    let _x = 1; 
    let _y = 2;

    let s = Some(String::from("hello"));

    // if let Some(_c) = s {   // 只忽略变量， 依然会发生所有权转移
    // // if let Some(c) = s {
    //     println!("found a string");
    // }

    // println!("s: {:?}", s);

    if let Some(_) = s {   // 忽略变量， 不会发生所有权转移
        println!("found a string");
    }

    println!("s: {:?}", s);
}