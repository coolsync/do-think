// 2| drop 提前释放
struct Dog {
    name: String,
}

impl Drop for Dog {
    fn drop(&mut self) {
        println!("{} leave.", self.name);
    }
}

// rust provide std::mem::drop(_x: T) method
fn main() {
    let a = Dog{name: String::from("wangcai")};
    let b = Dog{name: String::from("dahuang")};
    
    drop(b);
    drop(a);

    println!("+++++++++++++++++++");
    println!("Hello, world!");
}
