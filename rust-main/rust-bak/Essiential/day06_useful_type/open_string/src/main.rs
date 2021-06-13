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
