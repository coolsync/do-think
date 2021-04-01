
```bash
Compiling open_hashmap v0.1.0 (/home/dart/DoThinking/RustMain/Essiential/day06_hashmap/open_hashmap)
warning: variable does not need to be mutable
  --> src/main.rs:60:9
   |
60 |     let mut text = "hello world yes world";
   |         ----^^^^
   |         |
   |         help: remove this `mut`
   |
   = note: `#[warn(unused_mut)]` on by default

warning: 1 warning emitted

    Finished dev [unoptimized + debuginfo] target(s) in 0.62s
     Running `target/debug/open_hashmap`
v = 10
```



```bash
   Compiling open_err2 v0.1.0 (/home/dart/DoThinking/RustMain/Essiential/day08_mod_package/open_err2)
error[E0308]: mismatched types
  --> src/main.rs:14:33
   |
14 | fn read_username_from_file() -> Result<String, io::Error> {
   |    -----------------------      ^^^^^^^^^^^^^^^^^^^^^^^^^ expected enum `std::result::Result`, found `()`
   |    |
   |    implicitly returns `()` as its body has no tail or `return` expression
...
25 |     };
   |      - help: consider removing this semicolon
   |
   = note:   expected enum `std::result::Result<String, std::io::Error>`
           found unit type `()`

error: aborting due to previous error

For more information about this error, try `rustc --explain E0308`.
error: could not compile `open_err2`

To learn more, run the command again with --verbose.

```

solution:	将最后的分号去掉

```rust
fn read_username_from_file() -> Result<String, io::Error> {
    let f = File::open("h.txt");
    let mut f = match f {
        Ok(f) => f,
        Err(e) => return Err(e),
    };

    let mut s = String::new();
    match f.read_to_string(&mut s) {
        Ok(_) => Ok(s),
        Err(e) => Err(e),
    };
}
```



```rust
fn read_username_from_file() -> Result<String, io::Error> {
    let f = File::open("h.txt");
    let mut f = match f {
        Ok(f) => f,
        Err(e) => return Err(e),
    };

    let mut s = String::new();
    match f.read_to_string(&mut s) {
        Ok(_) => Ok(s),
        Err(e) => Err(e),
    }
}
```

