# Day07



## 1 | use mod



cargo new open_crate2

cargo new --lib mylib



mylib	--> factory.rs

```rust
pub mod produce_refrigerator {  // 生产冰箱
    pub fn produce_re() {
        println!("produce refrigerator!");
    }
}

pub mod produce_washing_machines {  // 生产洗衣机   // 私有 不可调用
    pub fn produce_washing() {
        println!("produce washing machine!");
    }

    pub fn produce_re() {
        println!("produce washing machine!");
    }
}
```



mylib --> lib.rs

```rust
pub mod factory;
#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}
```



指定 path

```bash
[package]
name = "open_crate2"
version = "0.1.0"
authors = ["dart"]
edition = "2018"

# See more keys and their definitions at https://doc.rust-lang.org/cargo/reference/manifest.html

[dependencies]
mylib = {path = "./mylib"}
```



src/main.rs

```rust
// use mylib::factory::produce_refrigerator;
// use mylib::factory::produce_refrigerator::produce_re;
// use mylib::factory::produce_washing_machines;
// use mylib::factory::produce_washing_machines as W;

use mylib::factory::*;

fn main() {
    mylib::factory::produce_refrigerator::produce_re(); // 绝对路径
    produce_refrigerator::produce_re(); // 使用 use， 推荐做法, path定位到 调用的上一级
    // produce_re();

    produce_washing_machines::produce_washing();
    produce_washing_machines::produce_re();
    
    // W::produce_washing();
    
    println!("Hello, world!");
}
```





# day08

## 2 | struct and mod



```rust
mod mod_a {
    #[derive(Debug)]    // 能够 print struct ino
    pub struct A {
        pub number: i32,
        name: String,
    }

    impl A {
        pub fn new_a() -> A {
            A {
                number: 1,
                name: String::from("A"),
            }
        }

        pub fn get_a(&self) {
            println!("number: {}, name: {}", self.number, self.name);
        }
    }

    pub mod mod_b {
        pub fn print_B() {
            println!("B");
        }

        pub mod mod_c {
            pub fn print_C() {
                println!("C");
                super::print_B();   // call Up-level function 
            }
        }
    }
}

// use mod_a::A;
use mod_a::A as A1;

fn main() {
    // let a = mod_a::A::new_a();    // 绝对 path
    // let a = A::new_a(); // 使用 use
    let a = A1::new_a();
    a.get_a();

    let number_a = a.number;
    // let name_a = a.name;

    println!("++++++++++++++");
    mod_a::mod_b::mod_c::print_C();
}

```



## 3 | crypto pkg usage



```rust
extern crate crypto;    // ref out pkg

use crypto::digest::Digest; // input_str
use crypto::sha3::Sha3;

fn main() {
    // get hash value
    let mut hasher = Sha3::sha3_512();
    hasher.input_str("hello world");
    let result = hasher.result_str();

    println!("hash: {}", result);
}
```

Cargo.toml:

```toml
[package]
name = "open_crate4_crypto"
version = "0.1.0"
authors = ["dart"]
edition = "2018"

# See more keys and their definitions at https://doc.rust-lang.org/cargo/reference/manifest.html

[dependencies]
rust-crypto = "^0.2"
```

## 4 | err



### 1 	可恢复 err, 不可恢复 err

（1）可恢复错误通常代表向用户报告错误和重试操作是合理的情况，例如未找到文件。rust中使用Result<T,E>来实现。

（2）不可恢复错误是bug的同义词，如尝试访问超过数组结尾的位置。rust中通过panic！来实现。



### 2 	pinic!



### 3	RUST_BACKTRACE	

`RUST_BACKTRACE=1`

`RUST_BACKTRACE=full`

run with `RUST_BACKTRACE=1` environment variable to display a backtrace

Some details are omitted, run with `RUST_BACKTRACE=full` for a verbose backtrace.



### 4	Result<T,E>



### 5 	Abbreviation（简写）



```rust
// 1| 可恢复error, 不可恢复error, two kind  
//（1）可恢复错误通常代表向用户报告错误和重试操作是合理的情况，例如未找到文件。rust中使用Result<T,E>来实现。
//（2）不可恢复错误是bug的同义词，如尝试访问超过数组结尾的位置。rust中通过panic！来实现。

// 2| panic!

// 3| 使用BACKTRACE=1

// 4| Result<T,E>

// Result<T,E> {
//     Ok(T),
//     Err(E),
// }

// 5| 简写

use std::fs::File;

fn main() {
    // let f = File::open("hello.txt");
    // let r = match f {
    //     Ok(file) => file,
    //     Err(e) => panic!("open file err: {:?}", e),
    // };

    // let f = File::open("hello.txt").unwrap();
    let f = File::open("hello.txt").expect("open file err");
    
    panic!("panic");
}
```





# day09



## test

cargo test



`/home/dart/DoThinking/RustMain/Essiential/day09_test/open_test/mylib/src/animal.rs`



```rust
pub mod dog {
    pub fn hello() {
        println!("wangwang");
    }

    pub fn is_dog() -> bool {
        true
    }
}

pub mod bird {
    pub fn hello() {
        println!("niaoniao");
    }

    pub fn is_bird() -> bool {
        true
    }
}
```



`/home/dart/DoThinking/RustMain/Essiential/day09_test/open_test/mylib/src/lib.rs`

```rust
pub mod animal;

#[cfg(test)]
mod tests {
    // use crate::animal::*;
    use crate::animal::dog;
    // use crate::animal::bird;
    
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
    #[test]
    fn use_dog() {
        assert_eq!(true, dog::is_dog())
    }

    #[test]
    fn use_bird() {
        assert_eq!(true, crate::animal::bird::is_bird())
    }
}
```

