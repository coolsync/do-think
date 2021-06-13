

# Day04

# 1 struct



```rust
fn main() {
    //1 define struct
    #[derive(Debug)] // print before, add

    struct User {
        name: String,
        count: String, // 账户
        nonce: u64,    // balance
        active: bool,  // 是否在线
    };

    //2 create struct instance
    let bob = User {
        name: String::from("bob"),
        count: String::from("80001000"),
        nonce: 10000,
        active: true,
    };

    println!("bob: {:?}", bob);
    println!("bob: {:#?}", bob);

    // 3 modify struct filed
    let mut paul = User {
        name: String::from("bob"),
        count: String::from("80001000"),
        nonce: 10000,
        active: true,
    };

    paul.nonce = 20000;

    // 4 params name and filed name on same name, simple method
    let name = String::from("jerry");
    let count = String::from("89007777");
    let nonce = 200000;
    let active = false;

    // let mark = User {
    //     name: name,
    //     count: count,
    //     nonce: nonce,
    //     active: active,
    // };

    let user1 = User {
        name,
        count,
        nonce,
        active,
    };

    // 5 from other struct create instance	
    let user2 = User {
        name: String::from("alice"),
        ..user1
    };

    println!("user2 name = {}", user2.name);
    println!("user2 count = {}", user2.count);

    // 6 tuple struct// A unit struct
    struct Point(char, i32, i32);

    let a = Point('1', 10, 20);
    let b = Point('2', 30, 5);

    println!("x = {}, y= {}", a.0, a.1);

    // 7 没有任何 field 的 类单元 struct, A unit struct
    // Unit structs, which are field-less, are useful for generics.
    struct A {};

    // 8 print struct filed info
}
```



## method



```rust
#[derive(Debug)]

struct Dog {
    name: String,
    age: u8,
    weight: f32,
}

// method
impl Dog {
    fn get_name(&self) -> &str {
        // &self.name[..]   // ?
        &(self.name[..])
    }
    fn get_age(&self) -> u8 {
        self.age
    }

    // fn get_weight(&self) -> f32 {
    //     self.weight
    // }

    fn show(&self) {
        println!("ho ho ho");
    }
}

// 分开method
impl Dog {
    fn get_weight(&self) -> f32 {
        self.weight
    }
}

fn main() {
    let dog = Dog {
        name: String::from("pangzhi"),
        age: 1,
        weight: 30.8,
    };

    println!("dog = {:#?}", dog);

    // Dog::show();
    dog.show();
    println!("name = {}", dog.get_name());
    println!("age = {}", dog.get_age());
    println!("weight = {}", dog.get_weight());
}
```



# Day05

# 2 Enum

```rust
// 1 c-like 方式
enum IPAddrKind {
    V4,
    V6,
}

struct IPAddr {
    kind: IPAddrKind,
    address: String,
}

// 2 r recommand 方式
enum IPAddr2 {
    V4(String),
    V6(String),
}

// 3 可以是 不同的 type
enum IPAddr3 {
    V4(u8, u8, u8, u8),
    V6(String),
}

// 4 经典用法
enum Message {
    Quit,   // unit struct
    Move{x: i32, y: i32}, // struct
    Change(i32, i32, i32),  // tuple
    Write(String)   // pass &str
}

// 5 enum type method, match (like switch)
impl Message {
    fn print(&self) {
        match *self {
            Message::Quit => println!("Quit"),
            Message::Move{x, y} => println!("Move x: {}, y: {}", x, y),
            Message::Change(a, b, c) => println!("Change a: {}, b: {}, c: {}", a, b, c),
            
            // Message::Write(&s) => println!("Write: {}", s),
            // _ => println!("Write"), // defult
            Message::Write(_) => println!("Write"),
        }
    }
}

fn main() {
    
    // 1
    let i1 = IPAddr {
        kind: IPAddrKind::V4,
        address: String::from("127.0.0.1"),
    };
    let i2 = IPAddr {
        kind: IPAddrKind::V6,
        address: String::from("::1"),
    };

    
    // 2
    let i1 = IPAddr2::V4(String::from("127.0.0.1"));
    let i2 = IPAddr2::V6(String::from("::1"));


    // 3
    let i1 = IPAddr3::V4(127, 0, 0, 1);
    let i2 = IPAddr3::V6(String::from("::1"));

   
    // 4
    let quit = Message::Quit;
    quit.print();

    let mv = Message::Move{x: 1, y: 2};
    mv.print();

    let change = Message::Change(10, 20, 30);
    change.print();

    let wri = Message::Write(String::from("hello"));
    wri.print();
}
```





## Option build-in

handler nil pointer.



```rust
// 1 Option 是 std 内 一个 enum, like:
// enum Option<T> {
//     Some(T),
//     None
// }

// handler nil pointer.

// 2 usage 方式

fn main() {
    let some_num = Some(1);
    let some_str = Some(String::from("a string"));
    let absent_num: Option<i32> = None;
    let x: i32 = 5;
    // let y = Some(5);
    let y: Option<i32> = Some(5);

    // let sum = x + y;    // no implementation for `i32 + std::option::Option<{integer}>`

    let mut temp = 0;
    match y {
        Some(y) => {
            temp = y;
        }
        None => {
            println!("do nothing")
        }
    }

    // let result = plus_one(y);
    // match result {
    //     Some(i) => println!("reuslt: {}", i),
    //     None => println!("nothing"),
    // }
    
    if let Some(value) = plus_one(y) {
        println!("value: {}", value);
    } else {
        println!("do nothing");
    }

    let sum = x + temp;
    println!("sum = {}", sum);
}

fn plus_one(i: Option<i32>) -> Option<i32> {
    match i {
        Some(i) => Some(i + 1),
        None => None,
    }
}
```



# Day06

# 3 Vector



```rust
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

    v3.push(6); // mutable borrow occurs here
    
    println!("first = {}", first); // immutable borrow later used here
}
```





# 4 String



```rust
// 1 create a 空的 string

// 2 via 字面值 create string
//  2.1 String::from 方式
//  2.2 str 方式

// 3 update String
//  3.1 push_str
//  3.2 push
//  3.3 使用 "+" 合并 string
//  3.4 format!

// 4 String index

// 5 str index

// 6 遍历
//      chars
//      bytes
fn main() {
    // 1
    let mut s0 = String::new(); // std::string::String
    s0.push_str("hello");
    println!("s0 = {}", s0);

    // 2.1
    let s1 = String::from("init a string");
    println!("s1 = {}", s1);

    // 2.2
    let s1 = "init a string".to_string();
    println!("s1 = {}", s1);

    // 3.1
    let mut s2 = String::from("hello");
    s2.push_str(" world");
    let ss = " !".to_string();
    s2.push_str(&ss); // 此处 为 ref, ss 依然可打印
    println!("s2 = {}", s2);
    println!("ss = {}", ss);

    // 3.2
    let mut s2 = String::from("tea");
    s2.push('m');
    // s2.push('mx'); err character literal may only contain one codepoint
    // s2.push("x"); err
    println!("s2 = {}", s2);

    // 3.3
    let s1 = "hello".to_string();
    let s2 = String::from(" world");
    let s3 = s1 + &s2; // s1 method ownership move to s3; s2 looking a function ref
    println!("s3 = {}", s3);

    // move occurs because `s1` has type `String`, which does not implement the `Copy`trait
    // println!("s1 = {}", s1); err // value borrowed here after move
    println!("s2 = {}", s2);

    // 3.4
    let s341 = String::from("tic");
    let s342 = String::from("tac");
    let s343 = String::from("toe");
    let s344 = format!("{}-{}-{}", s341, s342, s343); //format!和println!类似
    println!("s344 = {}", s344);
    println!("s341 = {}", s341);
    println!("s342 = {}", s342);
    println!("s343 = {}", s343);

    // 4
    let s4 = String::from("hello");
    // let s41 = s4[1];
    println!("s4.len = {}   ", s4.len());

    let s4 = String::from("你好");
    println!("s4.len = {}", s4.len());
    // let s41 = s4[1];

    // 5
    let hello = "你好";
    let h5 = &hello[0..3];
    println!("h5 = {}", h5);

    // let h6 = &hello[0..2];
    // println!("h6 = {}", h6);

    // 6
    for c in s4.chars() {
        println!("c = {}", c);
    }

    println!("+++++++++++++++++++");
    for b in s4.bytes() {
        println!("b = {}", b);
    }
    println!("+++++++++++++++++++");
}
```



# Day07



# HashMap

```rust
// 1| HashMap<K, V>
// 2| 创建 HashMap
// 3| 读取
// 4| 遍历
// 5| 更新

use std::collections::HashMap;

fn main() {
    // 1 HashMap<K,V>
    
    // 2 create HashMap
    let mut scores: HashMap<String, i32> = HashMap::new();
    scores.insert(String::from("blue"), 10);
    scores.insert(String::from("red"), 20);

    let keys: Vec<String> = vec![String::from("blue"), String::from("red")];
    let values = vec![10, 20];

    let scores: HashMap<_,_> = keys.iter().zip(values.iter()).collect();

    // 3 read
    let k = String::from("blue");
    if let Some(v) = scores.get(&k) {   // get 返回的是一个 Option enum
        println!("v = {}", v);
    }


    let k = String::from("yellow");
    let v = scores.get(&k);
    match v {
        Some(value) => println!("value: {}", value),
        None => println!("None"),
    }

    // 4 traverse： 会以任意的顺序 遍历出来
    println!("+++++++++++");
    for (key, value) in &scores {
        println!("key: {}, value: {}", key, value);
    }
    println!("+++++++++++");

    // 5 update
    // 直接插入
    let mut ss: HashMap<String, i32>= HashMap::new();
    ss.insert(String::from("one"), 1);
    ss.insert(String::from("two"), 2);
    ss.insert(String::from("three"), 3);
    ss.insert(String::from("one"), 3);
    println!("ss = {:?}", ss);

    // 健不存在的时候才插入
    let mut ss1: HashMap<String, i32>= HashMap::new();
    ss1.insert(String::from("one"), 1);
    ss1.insert(String::from("two"), 2);
    ss1.insert(String::from("three"), 3);
    ss1.entry(String::from("one")).or_insert(3);
    println!("ss1 = {:?}", ss1);

    // due to old val update a new val
    let text = "hello world yes world";
    let mut map = HashMap::new();

    for word in text.split_whitespace() {
        let count = map.entry(word).or_insert(0);
        *count += 1;    // map{"hello": 1}
    }

    println!("map = {:?}", map);
}
```

