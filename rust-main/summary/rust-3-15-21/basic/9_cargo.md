# day16



# Use Cargo



## 1 优化 level



```toml
[package]
name = "open_cargo1"
version = "0.1.0"
authors = ["dart"]
edition = "2018"

# See more keys and their definitions at https://doc.rust-lang.org/cargo/reference/manifest.html

[dependencies]
[profile.dev]
opt-level = 0

[profile.release]
opt-level = 3
```



```bash
[dart@localhost open_cargo1]$ cargo build
   Compiling open_cargo1 v0.1.0 (/home/dart/DoThinking/RustMain/Advanced/day16_cargo/open_cargo1)
    Finished dev [unoptimized + debuginfo] target(s) in 0.40s

[dart@localhost open_cargo1]$ cargo build
   Compiling open_cargo1 v0.1.0 (/home/dart/DoThinking/RustMain/Advanced/day16_cargo/open_cargo1)
    Finished dev [optimized + debuginfo] target(s) in 0.77s


[profile.release]
opt-level = 3
[dart@localhost open_cargo1]$ cargo build --release
   Compiling open_cargo1 v0.1.0 (/home/dart/DoThinking/RustMain/Advanced/day16_cargo/open_cargo1)
    Finished release [unoptimized] target(s) in 0.74s

[profile.release]
opt-level = 0
[dart@localhost open_cargo1]$ cargo build --release
   Compiling open_cargo1 v0.1.0 (/home/dart/DoThinking/RustMain/Advanced/day16_cargo/open_cargo1)
    Finished release [optimized] target(s) in 0.66s


```



## 2 doc comment

/src/lib.rs

```rust
//! # My Crate
//! 
//! `my_crate` is a collection of utilities to make performing certain
//! calculations more convenient.

/// And one to the number given.
/// 
/// #Example
/// ```
/// let five = 5;
/// assert_eq!(6, my_crate::add_one(five));
/// ```

pub fn add_one(x: i32) -> i32 {
    x + 1
}

// #[cfg(test)]
// mod tests {
//     #[test]
//     fn it_works() {
//         assert_eq!(2 + 2, 4);
//     }
// }
```



command:



```bash
cargo doc

cargo doc --open

cargo test
```



process bug:

```bash
test result: ok. 0 passed; 0 failed; 0 ignored; 0 measured; 0 filtered out; finished in 0.00s

   Doc-tests my_crate

running 1 test
test src/lib.rs - add_one (line 4) ... FAILED

failures:

---- src/lib.rs - add_one (line 4) stdout ----
error[E0425]: cannot find function `add_one` in this scope
 --> src/lib.rs:6:15
  |
4 | assert_eq!(6, add_one(five));
  |               ^^^^^^^ not found in this scope
  |
help: consider importing this function
  |
2 | use my_crate::add_one;
  |

error: aborting due to previous error

For more information about this error, try `rustc --explain E0425`.
Couldn't compile the test.

failures:
    src/lib.rs - add_one (line 4)

test result: FAILED. 0 passed; 1 failed; 0 ignored; 0 measured; 0 filtered out; finished in 0.05s

error: test failed, to rerun pass '--doc'
```



## 3 release to cargo.io



## 4 workspace

### [Creating a Workspace](https://doc.rust-lang.org/nightly/book/ch14-03-cargo-workspaces.html#creating-a-workspace)

### [Creating the Second Package in the Workspace](https://doc.rust-lang.org/nightly/book/ch14-03-cargo-workspaces.html#creating-the-second-package-in-the-workspace)



```bash
 cargo build
 cargo run -p adder
 cargo run -p adder2
```



