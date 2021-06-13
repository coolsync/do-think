fn main() {
    println!("Hello, world!");
    // 1. Function Parameters
    f1(5, 6);

    // 2. Function Bodies Contain Statements and Expressions
    let a = 5;
    let b = {
        let a = 6;
        a + 1
    };
    println!("a = {}", a);
    println!("b = {}", b);

    // 3. Functions with Return Values
    let a2 = f2(5);
    println!("a2 = {}", a2);
    
}

fn f1(x: i32, y: i32) {
    println!("x = {}, y = {}", x, y)
}

fn f2(x: i32) -> i32 {
    x + 1
}