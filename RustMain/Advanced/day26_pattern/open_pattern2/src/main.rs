// 1| match literals (字面值)
// fn main() {
//     let x = 1;

//     match x {
//         1 => println!("one"),
//         2 => println!("two"),
//         _ => println!("other"),    
//     }
// }

// 2| match named variables (命名变量)
// fn main() {
//     let x = Some(5);
//     let y = 10; // 位置1

//     match x {
//         Some(50) => println!("50"),
//         Some(y) => println!("y: {:?}", y), // 此处是 位置2
//         _ => println!("default case x = {:?}", x),
//     }

//     println!("x: {:?}, y: {:?}", x, y); // 此处是 位置1
// }

// 3| match 多个模式
// fn main() {
//     let x = 1;
//     match x {
//         1|2 => println!("1 or 2"),  // | 表示： match 1 or 2
//         3 => println!("3"),
//         _ => println!("other"),
//     };
// }

// 4| 通过 .. match
fn main() {
    // let x = 2;
    
    // match x {
    //     1..=5 => println!("1 to 5"),    // 1|2|3|4|5 => println!("1 to 5")
    //     _ => println!("ohter"),
    // };

    let x = 'c';
    
    match x {
        'a'..='j' => println!("1"),
        'k'..='z' => println!("2"),
        _ => println!("other"),
    }
}
