// vector is a dynamic array that can store any type,
// can increase and compress data
// vector是一个能够存放任意类型的动态数组，能够增加和压缩数据

// 1 create 空的 vector: Vec<T>， 没有实际意义
// 2 create 包含 初始值的 vector
// 3 drop vector
// 4 读取元素 read element
// 5 update
// 6 遍历
// 7 use enum

fn main() {
    // 1
    let v: Vec<i32> = Vec::new();
    // let mut v: Vec<i32> = Vec::new();
    // v.push(1);

    // 2
    let v = vec![1, 2, 3];

    // 3
    {
        let v1 = vec![1, 2, 3];
    }

    // 4
    let one = &v[0];
    // let four = &v[4];
    println!("one = {}", *one);
    println!("one = {}", one); // auto conv

    // (2) recommend method:
    // match v.get(1) {
    match v.get(4) {
        Some(value) => println!("value = {}", value),
        _ => println!("None"),
    }

    // 5
    let mut v2: Vec<i32> = Vec::new();
    v2.push(1);
    v2.push(2);
    v2.push(3);

    // 6
    // 不可变 遍历
    for i in &v2 {
        println!("i = {}", i)
    }

    // 可变 遍历
    for i in &mut v2 {
        *i += 1;
        println!("i = {}", i)
    }

    // 7
    enum Context {
        Text(String),
        Float(f32),
        Int(i32),
    };

    let c = vec![
        Context::Text(String::from("string")),
        Context::Float(0.001),
        Context::Int(1),
    ];

    // 8 补充
    let mut v3 = vec![1, 2, 3, 4, 5];
    let first = &v3[0]; // -- immutable borrow occurs here  只能 借用 一次， v3 被锁定 immutable

    // v3.push(6); // mutable borrow occurs here
    
    println!("first = {}", first); // immutable borrow later used here
}
