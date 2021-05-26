//1、在任意给定时间，有了可变引用之后不能再有不可变引用
//2、引用必须有效

fn main() {
    let r = dangle();
    println!("Hello, world!");
}

fn dangle() -> &String {
    let s = String::from("hello");
    &s
}

/* 
error[E0106]: missing lifetime specifier
 --> main.rs:6:16
  |
6 | fn dangle() -> &String {
  |                ^ expected named lifetime parameter
  |
  = help: this function's return type contains a borrowed value, but there is no value for it to be borrowed from
help: consider using the `'static` lifetime
  |
6 | fn dangle() -> &'static String {
  |                ^^^^^^^^

error: aborting due to previous error
*/