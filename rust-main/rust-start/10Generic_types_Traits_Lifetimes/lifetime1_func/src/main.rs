// fn largest(x: &str, y: &str) -> &str {
// fn largest<'a>(x: &'a str, y: &'a str) -> &'a str {
fn largest<'c>(x: &'c str, y: &'c str) -> &'c str {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

fn get_str<'a>(x: &'a str, y: &str) -> &'a str {
    x
}

// error:
// fn return_str<'a>(x: &'a str, y: &'a str) -> &'a str {
//     let ss = String::from("hello");
//     ss.as_str() // returns a value referencing data owned by the current function
// }

fn main() {
    let s1 = String::from("abcd");
    let s2 = "xy";

    let r = largest(s1.as_str(), s2);
    println!("r: {}", r);
    
    let r2 = get_str(s1.as_str(), s2);
    // let r3 = return_str(s1.as_str(), s2);

    println!("Hello, world!");
}

//   = help: this function's return type contains a borrowed value, but the signature does not say whether it is borrowed from `x` or `y`
// help: consider introducing a named lifetime parameter