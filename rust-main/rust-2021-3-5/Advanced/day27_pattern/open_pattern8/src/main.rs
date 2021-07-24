fn main() {
    let numbers = (1, 2, 3, 4, 5, 6, 7);
    match numbers {
        (first, .., last) => println!("first: {}, last: {}", first, last),
    }

    // error: `..` can only be used once per tuple pattern
    // match numbers {
    //     (.., second, ..) => println!("second: {}", second),
    // }
    
    println!("Hello, world!");
}
