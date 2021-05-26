
fn takes_ownership(something: String) -> String {
    println!("{}", something); 
    something   // scope in main fn
} 

fn make_cp(i :i32) {
    println!("{}", i);
}
fn main() {
    let s = String::from("nice");   // heap
    // takes_ownership(s); // invalid, scope in takes_ownership fn
    // println!("{}", s); 

    let s1 = takes_ownership(s);
    println!("{}", s1);

    let i = 1;  // stack
    make_cp(i);
    println!("i = {}", i);  // valid
}
