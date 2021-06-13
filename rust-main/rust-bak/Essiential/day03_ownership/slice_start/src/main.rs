//1、字符串slice是String中一部分值的引用
//2、字面值就是slice
//3、其它类型slice

fn main() {
    let s = String::from("hello world");

    let h = &s[0..5];
    let h = &s[0..=4];
    let h = &s[..=4];
    let h = &s[..5];
    println!("h = {}", h);

    let w = &s[6..11];
    let w = &s[6..];
    let w = &s[6..=10];
    let w = &s[..];
    println!("w = {}", w);

    let ss = String::from("你好");

    // let nihao = &ss[0..1];
    // thread 'main' panicked at 'byte index 1 is not a char boundary; it is inside '你' (bytes 0..3) of `你好`'
    // println!("nihao = {}", nihao);

    // let s1 = "hhh";    // &str
    // let h1 = s1[0..1];
    // println!("h1 = {}", h1);  // the type `str` cannot be indexed by `{integer}` string indices are ranges of `usize`

    let a = [1, 2, 3, 4, 5];

    let aa = &a[1..3];  // 2, 3
    println!("aa = {}", aa[0]);
    println!("aa = {}", aa[1]);
}
