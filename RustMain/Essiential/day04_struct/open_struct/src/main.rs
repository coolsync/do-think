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
