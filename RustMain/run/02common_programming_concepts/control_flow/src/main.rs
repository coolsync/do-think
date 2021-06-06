fn main() {
    // use if in let statement
    let condition = true;

    let num = if condition {
        5
    } else {
        6
        // "six"   // `if` and `else` have incompatible types, 
    };

    println!("num = {}", num);

    println!("Hello, world!");
}
