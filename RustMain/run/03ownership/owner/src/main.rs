// 1  rust 通过 ownership 机制来管理内存，编译器在编译时由 ownership 机制对内存使用进行check

// 2  heap and stack

//   compile time data type space  fixed, allocation to stack

//   compile time, data type space not fixed, allocation to heap

// 3  scope 作用域

// 4  String 内存回收

// 5  move

// 6  clone

// 7  stack data copy

// 8  function and scope

fn main() {
    // 3  scope 作用域
    let x = 1;
    {
        let y = x;
        println!("x = {}", x); // stack data auto impl Copy trait
        println!("y = {}", y);
    } // leave this scope, y auto

    // println!("y = {}", y);  // cannot find value `y` in this scope

    // 4  String 内存回收
    {
        let mut s1 = String::from("hello");
        s1.push_str(" world");
        println!("s1 = {}", s1);

        let s2 = s1; // String type s1 leave, call drop method
                     // println!("s1 = {}", s1); //  err: value brrowed
        println!("s2 = {}", s2);

        let s3 = s2.clone(); // impl Copy trait, s2, s3 point to same heap addr
        println!("s2 = {}", s2);
        println!("s3 = {}", s3);
    }

    // impl Copy type
    // All the integer types, such as u32.
    // The Boolean type, bool, with values true and false.
    // All the floating point types, such as f64.
    // The character type, char.
    // Tuples, if they only contain types that also implement Copy. For example, (i32, i32) implements Copy, but (i32, String) does not.

    println!("Hello, world!");
}
