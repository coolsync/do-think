// 引用与借用
// 获取值的所有权
// fn main() {
//     let s1 = String::from("hello");

//     let (s2, len) = calculate_length(s1);

//     println!("The length of '{}' is {}.", s2, len);
// }

// fn calculate_length(s: String) -> (String, usize) {
//     let length = s.len(); // len() 返回字符串的长度

//     (s, length)
// }

// 它以一个对象的引用作为参数而不是获取值的所有权：
// fn main() {
//     let s1 = String::from("hello");

//     let len = calculate_length(&s1);

//     println!("The length of '{}' is {}.", s1, len);
// }

// fn calculate_length(s: &String) -> usize {
//     s.len()
// }

// 可变引用
// 1. 正如变量默认是不可变的，引用也一样。（默认）不允许修改引用的值。
// fn main() {
//     let mut s = String::from("hello");

//     change(&mut s);
// }

// fn change(some_string: &mut String) {
//     some_string.push_str(", world");
// }

// fn main() {
// 2. 不过可变引用有一个很大的限制：在特定作用域中的特定数据只能有一个可变引用
//     let mut s = String::from("hello");

//     let r1 = &mut s;
//     let r2 = &mut s;    // cannot borrow `s` as mutable more than once at a time

//     println!("{}, {}", r1, r2);

// 3. 可以使用大括号来创建一个新的作用域，以允许拥有多个可变引用，只是不能 同时 拥有

// let mut s = String::from("hello");

// {
//     let r1 = &mut s;
// } // r1 在这里离开了作用域，所以我们完全可以创建一个新的引用

// let r2 = &mut s;

// 4. 类似的规则也存在于同时使用可变与不可变引用中。这些代码会导致一个错误
// let mut s = String::from("hello");

// let r1 = &s; // 没问题
// let r2 = &s; // 没问题
// let r3 = &mut s; // 大问题  // cannot borrow `s` as mutable because it is also borrowed as immutable

// println!("{}, {}, and {}", r1, r2, r3);

// solution:
// let mut s = String::from("hello");

// let r1 = &s; // 没问题
// let r2 = &s; // 没问题
// println!("{} and {}", r1, r2);
// // 此位置之后 r1 和 r2 不再使用

// let r3 = &mut s; // 没问题
// println!("{}", r3);
// }

// 悬垂引用（Dangling References）
// 在具有指针的语言中，很容易通过释放内存时保留指向它的指针而错误地生成一个 悬垂指针（dangling pointer），所谓悬垂指针是其指向的内存可能已经被分配给其它持有者。I

fn main() {
    // let reference_to_nothing = dangle();
    let reference_to_nothing = no_dangle();

}

// this function's return type contains a borrowed value, but there is no value for it to be borrowed from
// fn dangle() -> &String { // dangle 返回一个字符串的引用
    
//     let s = String::from("hello"); // s 是一个新字符串
    
//     &s // 返回字符串 s 的引用
// } // 这里 s 离开作用域并被丢弃。其内存被释放。
// // 危险！
// 因为 s 是在 dangle 函数内创建的，当 dangle 的代码执行完毕后，s 将被释放。不过我们尝试返回它的引用。这意味着这个引用会指向一个无效的 String

// The solution here is to return the String directly:
fn no_dangle() -> String {
    let s = String::from("hello");

    s   // 所有权被移动出去，所以没有值被释放。
}
