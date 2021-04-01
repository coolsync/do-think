// Variable life circle
fn main() {
    // Err example
    // let r;
    // {
    // let x = 5;
    // r = &x;
    // }
    // println!("r: {}", r);
    
    let r;

    let x = 5;
    r = &x;
    println!("r: {}", r);
}
