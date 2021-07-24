fn main() {
    // 1 c-like 方式
    enum IPAddrKind {
        V4,
        V6,
    };

    struct IPAddr {
        kind: IPAddrKind,
        address: String,
    };

    let i1 = IPAddr {
        kind: IPAddrKind::V4,
        address: String::from("127.0.0.1"),
    };
    let i2 = IPAddr {
        kind: IPAddrKind::V6,
        address: String::from("::1"),
    };

    
    // 2 r recommand 方式
    enum IPAddr2 {
        V4(String),
        V6(String),
    }

    let i1 = IPAddr2::V4(String::from("127.0.0.1"));
    let i2 = IPAddr2::V6(String::from("::1"));


    // 3 可以是 不同的 type
    enum IPAddr3 {
        V4(u8, u8, u8, u8),
        V6(String),
    }

    let i1 = IPAddr3::V4(127, 0, 0, 1);
    let i2 = IPAddr3::V6(String::from("::1"));
}
