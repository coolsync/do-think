fn takes_ownership(some_string: String) -> String {
    println!("some_string: {}", some_string);
    some_string
}

fn make_copy(i: i32) {
    println!("i = {}", i)
}



fn main() {
    let s1 = String::from("hello");
    let s2 = takes_ownership(s1);   // s1 call drop method
    // println!("s1 = {}", s1);    // -- move occurs because `s1` has type `String`, which does not implement the `Copy` trait
    println!("s2 = {}", s2);
    
    let x = 1;
    make_copy(x);
    println!("x = {}", x); // stack data impl Copy Trait
    println!("Hello, world!");


    // Return Values and Scope
    let s3 = String::from("hello2");
    let (s4, len) = caculate_len(s3);
    println!("'{}' length is {}", s4, len)
}

fn caculate_len(s: String) -> (String, usize) {
    let len = s.len();
    (s, len)
}

