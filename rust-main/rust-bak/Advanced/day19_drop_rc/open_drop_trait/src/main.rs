// 1| Drop trait like other lang 析构 function，
// 当值离开作用域的时候执行此func的代码。
struct Dog {
    name: String,
}

impl Drop for Dog {
    fn drop(&mut self) {
        println!("Dog {} leave", self.name);
    }
}

fn main() {
    let a = Dog{name: String::from("xiaopangzi")};
    {
        let b = Dog{name: String::from("dahuang")};
        println!("0 ++++++++++++++++");    
    }   // call drop

    println!("1 ++++++++++++++++");

    println!("Hello, world!");
} // call drop
