use mac;

fn main() {
    let v = mac::my_vec![1,2,3];
    println!("v: {:?}", v);
    
    // mac::my_vec![1,2,3] 等价于
    // let mut tmp_vec = Vec::new();
    // tmp_vec.push(1);
    // tmp_vec.push(2);
    // tmp_vec.push(3);
    // tmp_vec
    
    println!("Hello, world!");
}
