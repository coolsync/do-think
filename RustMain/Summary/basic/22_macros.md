# day31

## macros1 declare



./Cargo.toml:

```shell
[workspace]
members = [
    "mac",
    "main",
]
```



./mac/src/lib.rc:

```rust
#[macro_export]
macro_rules! my_vec {
    ($($x: expr), *) => {	// match 0 or 1 次
        {
            let mut tmp_vec = Vec::new();
            $(
                tmp_vec.push($x);
            )*
            tmp_vec
        }
    };
}
```



./main/Cargo.toml:

```shell
[package]
name = "main"
version = "0.1.0"
authors = ["dart"]
edition = "2018"

# See more keys and their definitions at https://doc.rust-lang.org/cargo/reference/manifest.html

[dependencies]
mac = {path = "../mac"}
```



./main/src/main.rc:

```rust
use mac;

fn main() {
    let v = mac::my_vec![1,2,3];
    println!("v: {:?}", v);
    
    // mac::my_vec![1,2,3] 等价于
    // let mut tmp_vec = Vec::new();
    // tmp_vec.push(1);
    // tmp_vec.push(2);
    // tmp_vec.push(3);
    // tmp_vec
    
    println!("Hello, world!");
}
```





## macros2	[How to Write a Custom `derive` Macro](https://doc.rust-lang.org/book/ch19-06-macros.html#how-to-write-a-custom-derive-macro)

4、过程宏介绍
过程宏接收 Rust 代码作为输入，在这些代码上进行操作，然后产生另一些代码作为输出，而非像声明式宏那样匹配对应模式然后以另一部分代码替换当前代码。

定义过程宏的函数接受一个 TokenStream 作为输入并产生一个 TokenStream 作为输出。这也就是宏的核心：宏所处理的源代码组成了输入 TokenStream，同时宏生成的代码是输出 TokenStream。如下：
use proc_macro;
#[some_attribute]
pub fn some_name(input: TokenStream) -> TokenStream {
}

过程宏中的derive宏   fmt::Display trait
#[derive(Debug)]
struct A {
	a : i32,
}

说明：在hello_macro_derive函数的实现中，syn 中的 parse_derive_input 函数获取一个 TokenStream 并返回一个表示解析出 Rust 代码的 DeriveInput 结构体（对应代码syn::parse(input).unwrap();）。该结构体相关的内容大体如下：


    DeriveInput {
        // --snip--
        
    ident: Ident {
        ident: "Pancakes",
        span: #0 bytes(95..103)
    },
    data: Struct(
        DataStruct {
            struct_token: Struct,
            fields: Unit,
            semi_token: Some(
                Semi
            )
        }
    )
    }


tree:

```shell
├── hello_macro
│   ├── Cargo.lock
│   ├── Cargo.toml
│   ├── hello_macro_derive
│   │   ├── Cargo.lock
│   │   ├── Cargo.toml
│   │   ├── src
│   │   │   └── lib.rs
│   │   └── target
│   ├── src
│   │   └── lib.rs
│   └── target
│       ├── CACHEDIR.TAG
│       └── rls
│           ├── CACHEDIR.TAG
│           └── debug
├── main
│   ├── Cargo.lock
│   ├── Cargo.toml
│   ├── src
│   │   └── main.rs
│   └── target
└── Readme

```



./hello_macro/src/lib.rc

```rust
pub trait HelloMacro {
    fn hello_macro();
}
```

./hello_macro/hello_macro_derive/Cargo.toml:

```toml
[package]
name = "hello_macro_derive"
version = "0.1.0"
authors = ["dart"]
edition = "2018"

# See more keys and their definitions at https://doc.rust-lang.org/cargo/reference/manifest.html

[lib]
proc-macro = true

[dependencies]
syn = "1.0"
quote = "1.0"
```



./hello_macro/hello_macro_derive/src/lib.rc

```rust
extern crate proc_macro;
use crate::proc_macro::TokenStream;
use quote::quote;
use syn;

fn impl_hello_macro(ast: &syn::DeriveInput) -> TokenStream {
    let name = &ast.ident; // get struct name
    let gen = quote! {
        impl HelloMacro for #name {
            fn hello_macro() {
                println!("hello, in my macro, my name is {}", stringify!(#name))
                // 根据需要，stringify函数可以允许定制一个复杂对象的特定属性如何被格式化。 
            }
        }
    };
    gen.into()
}

#[proc_macro_derive(HelloMacro)]
pub fn hello_macro_derive(input: TokenStream) -> TokenStream {
    let ast = syn::parse(input).unwrap();   // parse a struct
    impl_hello_macro(&ast)
}
```

./main/Cargo.toml

```toml
[package]
name = "main"
version = "0.1.0"
authors = ["dart"]
edition = "2018"

# See more keys and their definitions at https://doc.rust-lang.org/cargo/reference/manifest.html

[dependencies]
hello_macro = {path = "../hello_macro"}
hello_macro_derive = {path = "../hello_macro/hello_macro_derive"}
```

./main/src/main.rs:

```rust
use hello_macro::HelloMacro;
use hello_macro_derive::HelloMacro;

#[derive(HelloMacro)]
struct Main;

fn main() {
    Main::hello_macro();    
}
```







6、类属性宏
类属性宏与自定义派生宏相似，不同于为 derive 属性生成代码，它们允许你创建新的属性。

例子：
可以创建一个名为 route 的属性用于注解 web 应用程序框架（web application framework）的函数：
#[route(GET, "/")]
fn index() {

#[route] 属性将由框架本身定义为一个过程宏。其宏定义的函数签名看起来像这样：
#[proc_macro_attribute]
pub fn route(attr: TokenStream, item: TokenStream) -> TokenStream {

说明：类属性宏其它工作方式和自定义derive宏工作方式一致。

7、类函数宏
类函数宏定义看起来像函数调用的宏。类似于 macro_rules!，它们比函数更灵活。
例子：
如sql！宏，使用方式为：
let sql = sql!(SELECT * FROM posts WHERE id=1);
则其定义为：
#[proc_macro]
pub fn sql(input: TokenStream) -> TokenStream {

8、宏的资料推荐
https://danielkeep.github.io/tlborm/book/mbe-macro-rules.html









