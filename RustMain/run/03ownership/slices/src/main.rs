// Let’s think about the signature of this function:
// fn first_word(s: &String) -> ?

// 通过字节的字面值语法来寻找代表空格的字节, 如果找到了一个空格，返回它的位置。否则，使用 s.len() 返回字符串的长度：
// fn first_word(s: &String) -> usize {
//     let bytes = s.as_bytes(); // as_bytes 方法将 String 转化为字节数组

    // 创建一个迭代器
    // 从 .iter().enumerate() 中获取了集合元素的引用，所以模式中使用了 &
//     for (i, &item) in bytes.iter().enumerate() {
//         // we search for the byte that represents the space by using the byte literal syntax.
//         if item == b' ' {
//             return i;
//         }
//     }

//     s.len()
// }
// fn main() {
//     let mut s = String::from("hello world");

//     let word = first_word(&s); // word 的值为 5

//     s.clear(); // 这清空了字符串，使其等于 ""

    // word 在此处的值仍然是 5，
    // 但是没有更多的字符串让我们可以有效地应用数值 5。word 的值现在完全无效！
// }

// 1. 字符串 slice（string slice）是 String 中一部分值的引用
// 2. 字符串字面值就是 slice
// 3. 字符串 slice 作为参数
// 4. 其他类型的 slice
// fn main() {
//     let s = String::from("hello world");
//     let len = s.len();

//     let all = &s[..];
//     let all = &s[..len];
//     println!("all = {}", all);


//     let h = &s[0..5];
//     let h = &s[0..=4];
//     let h = &s[..=4];
//     println!("h = {}", h);

//     let w = &s[6..11];
//     let w = &s[6..len];
//     let w = &s[6..];
//     println!("w = {}", w);

//     let s1 = "Hello, world!"; // 这里 s 的类型是 &str：它是一个指向二进制程序特定位置的 slice
//     // 字符串字面值是不可变的；&str 是一个不可变引用。

//     let a = [1,2,3,4,5];
//     let slice = &a[..3];    
//     println!("slice = {:?}", slice); //  `[{integer}]` doesn't implement `std::fmt::Display`
    
// }

fn first_word(s: &String) -> &str {
    let bytes = s.as_bytes();   // Returns a byte slice of this String's contents.

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[..i];
        }
    }
    &s[..]
} 

fn first_word2(s: &str) -> &str {
    let bytes = s.as_bytes();   // Returns a byte slice of this String's contents.

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[..i];
        }
    }
    &s[..]
} 

fn main() {
    let s = String::from("hello world");

    let word = first_word(&s); 

    // 借用规则，当拥有某值的不可变引用时，就不能再获取一个可变引用。因为 clear 需要清空 String，它尝试获取一个可变引用。
    // s.clear(); // 错误! // cannot borrow `s` as mutable, as it is not declared as mutable

    println!("word: {}", word);


    let my_string = String::from("hello world");

    // first_word2 中传入 `String` 的 slice
    let word = first_word2(&my_string[..]); // reference `&str`

    let my_string_literal = "hello world";

    // first_word2 中传入字符串字面值的 slice
    let word = first_word2(&my_string_literal[..]);

    // 因为字符串字面值 **就是** 字符串 slice，
    // 这样写也可以，即不使用 slice 语法！
    let word = first_word2(my_string_literal);
}