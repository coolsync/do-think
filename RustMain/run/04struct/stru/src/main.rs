fn main() {
    // 1. define struct
    #[derive(Debug)]
    struct User {
        name: String,
        count: String, // 账户
        nonce: u64,    // balance
        active: bool,  // 是否在线
    }

    // 2. instance struct
    let user1 = User {
        name: String::from("bob"),
        count: String::from("80001000"),
        nonce: 10000,
        active: true,
    };

    println!("user1: {:?}", user1);
    println!("user1: {:#?}", user1);


    // 3. modify struct field value
    let mut user1 = User {
        name: String::from("bob"),
        count: String::from("80001000"),
        nonce: 10000,
        active: true,
    };

    user1.nonce = 20000;

    // 4. 变量与字段同名时的字段初始化简写语法
    let name = String::from("mark");
    let count = String::from("80001001");
    let nonce = 20000; 
    let active = false;
    
    // let user2 = User {
    //     name:name,
    //     count: count,
    //     nonce:nonce,
    //     active:active,
    // };
    let user2 = User {
        name,
        count,
        nonce,
        active,
    };

    // 5. 使用结构体更新语法从其他实例创建实例
    let user3 = User {
        name: String::from("paul"),
        nonce: 200000,
        ..user1
    };
    println!("user1 nonce = {}", user1.nonce);
    println!("user3 nonce = {}", user3.nonce);

    // 6. 使用没有命名字段的 tuple structs 来创建不同的类型
    //(1)字段没有名字
    //(2)圆括号
    struct Point(i32, i32);
    
    let a = Point(11, 22);
    let b = Point(99, 22);

    println!("a.0 = {}, a.1 = {}", a.0, a.1);

    // 7. 没有任何字段的类单元结构体, nil struct
    struct A{}

    // 8. print struct
    println!("Hello, world!");
}
