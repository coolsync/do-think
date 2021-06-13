fn main() {
    let mut s1 = String::from("hello");
    let mut s2 = "hello";
    println!("s1 == s2: {}", s1==s2);
    
    s1 = "doo".to_string();
    println!("s1 = {}", s1);

    let ss1 = &s1[0..2];
    s2 = "good";
    println!("s2 = {}", s2);
    println!("Hello, world!");
}
