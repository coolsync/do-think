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
    
    let i1 = IPAddr {
        kind: IPAddrKind::V4,
        address: String::from("127.0.0.1"),
    };
    let i2 = IPAddr {
        kind: IPAddrKind::V6,
        address: String::from("::1"),
    };

    
    let i1 = IPAddr2::V4(String::from("127.0.0.1"));
    let i2 = IPAddr2::V6(String::from("::1"));


    let i1 = IPAddr3::V4(127, 0, 0, 1);
    let i2 = IPAddr3::V6(String::from("::1"));

   
    let quit = Message::Quit;
    quit.print();

    let mv = Message::Move{x: 1, y: 2};
    mv.print();

    let change = Message::Change(10, 20, 30);
    change.print();

    let wri = Message::Write(String::from("hello"));
    wri.print();
}
