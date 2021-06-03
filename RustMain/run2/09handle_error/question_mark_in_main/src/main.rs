use std::fs::File;
use std::error::Error;
// ? 运算符可被用于返回 Result fn
// fn main() {
// let f = File::open("h.txt")?; // cannot use the `?` operator in a function that returns `()`
//     println!("Hello, world!");
// }

fn main() -> Result<(), Box<dyn Error>> {
    // let f = File::create("h.txt")?;
    let f = File::open("h.txt")?;

    Ok(())
}