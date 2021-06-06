/*
 * @Author: your name
 * @Date: 2021-06-05 16:53:16
 * @LastEditTime: 2021-06-05 17:02:00
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /match_coin/src/main.rs
*/

enum Coin {
    Penny,
    Nickel,
    Dime,
    Quarter,
}

fn main() {
    let some_u8_value = 0u8;
    match some_u8_value {
        1 => println!("one"),
        3 => println!("three"),
        5 => println!("five"),
        7 => println!("seven"),
        0 => println!("zero"),
        _ => (),
    }

    println!("Hello, world!");
}
