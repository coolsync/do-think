// Function life circle
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

fn main() {
    let s1 = String::from("abcd");
    let s2 = String::from("ab");
    // Extracts a string slice containing the entire `String`.
    let r = largest(s1.as_str(), s2.as_str());
    println!("r = {}", r);

    let ss = get_str(s1.as_str(), s2.as_str());
}
