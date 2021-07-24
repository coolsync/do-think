//@运算符允许我们在创建一个存放值的变量的同时，测试这个变量的值是否匹配模式。

enum Message {
    Hello { id: i32 },
}

fn main() {
    let msg = Message::Hello { id: 25 };

    match msg {
        Message::Hello { id: id_val @ 0..=9 } => println!("id_val: {}", id_val),

        Message::Hello { id: 10..=20 } => println!("large"),

        Message::Hello { id } => println!("id: {}", id),
    }
    println!("Hello, world!");
}

// fn main() {
//     let x = 5;

//     match x {
//         x @ 0..=9 => println!("x: {}", x),
//         _ => println!("other"),
//     }

//     let y = Some(5);

//     match y {
//         Some(y_val @ 0..=9) => println!("y_val: {}", y_val),
//         _ => println!("ohter")
//     }
// }