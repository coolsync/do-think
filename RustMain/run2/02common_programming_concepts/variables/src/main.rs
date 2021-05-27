
const MAX_PONITS:u32 = 100_00;

fn main() {
    // 1. Variables and variability
    let a = 5;
    println!("a = {}", a);

    let mut b: u32 = 6;
    println!("b = {}", b);

    b = 7;
    println!("b = {}", b);

    // 2. Shadowing
    let b: f32 = 1.1;
    println!("b = {}", b);

    // 3. Constants
    println!("MAX_POINTS: {}", MAX_PONITS);

    
    // let spaces = "   ";
    // spaces = spaces.len();
    // println!("spaces: {}", spaces);

    // tuple
    let tup:(i32, f64, i8) = (50, 6.4, 8);
    println!("tup: {:?}", &tup);
    // array
    let a = [3;5];
    println!("a = {:?}", a);

    println!("Hello, world!");


}