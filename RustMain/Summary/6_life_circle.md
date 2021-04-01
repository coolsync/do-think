# Life Circle



## 1 - Variable life



```rust
// Variable life circle
fn main() {
    // Err example
    // let r;
    // {
    // let x = 5;
    // r = &x;
    // }
    // println!("r: {}", r);
    
    let r;

    let x = 5;
    r = &x;
    println!("r: {}", r);
}
```



## 2 - Function life

```rust
// Function lifetime
// fn largest(x: &str, y: &str) -> &str {
// fn largest<'a>(x: &'a str, y: &'a str) -> &'a str {
fn largest<'c>(x: &'c str, y: &'c str) -> &'c str {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

fn get_str<'a>(x: &'a str, y: &str) -> &'a str {
    x
}

fn main() {
    let s1 = String::from("abcd");
    let s2 = String::from("ab");
    // Extracts a string slice containing the entire `String`.
    let r = largest(s1.as_str(), s2.as_str());
    println!("r = {}", r);

    let ss = get_str(s1.as_str(), s2.as_str());
}
```



## 3 - Struct life



```rust
// 1| in struct lifetime 
#[derive(Debug)]
struct A<'a> {
    name: &'a str,
}

fn main() {
    let n = String::from("paul");
    let a = A{name: &n};
    println!("a = {:#?}", a);
}
```

## 4 - Lifetime omit (省略)



（1）没有生命周期注解却能够编译，原因：早期的rust中必须显式的声明生命周期，后来rust团队将很明确的模式进行了注解的简化。



（2）遵守生命周期省略规则的情况下能明确变量的声明周期，则无需明确指定生命周期。

函数或者方法的参数的生命周期称为输入生命周期，而返回值的生命周期称为输出生命周期。



（3）编译器采用三条规则判断引用何时不需要生命周期注解，

当编译器检查完这三条规则后仍然不能计算出引用的生命周期，则会停止并生成错误。



（4）生命周期注解省略规则适用于fn定义以及impl块定义，如下：



a、每个引用的参数都有它自己的生命周期参数。例如如下：

​         一个引用参数的函数，其中有一个生命周期： fn foo<'a>(x: &'a i32)

​       两个引用参数的函数，则有两个生命周期 ：fn foo<'a, 'b>(x: &'a i32, y: &'b i32)

​        以此类推。



   b、如果只有一个输入生命周期参数，那么它被赋予所有输出生命周期参数：

​          fn foo(x: &i32) -> &i32   等价于  fn foo<'a>(x: &'a i32) -> &'a i32



​    c、如果方法有多个输入生命周期参数，不过其中之一因为方法的缘故为&self或者&mut self，

​       那么self的生命周期被赋予所有输出生命周期参数。例子在下面来看。

​    fn function(&self, x: &str, y: &str, ....) -> &str



```rust
// 2| lifetime 省略
fn get_a_str(s: &str) -> &str {
    s
}

fn main() {
    let n = String::from("paul");
    let a = A{name: &n};
    println!("a = {:#?}", a);

    let r = get_a_str(&n);
    println!("r = {}", r);
}
```



## 5 - Method lifetime



```rust
struct StrA<'a> {
    name: &'a str,
}

impl<'b> StrA<'b> {
    fn do_something(&self) -> u32 {
        3
    }

    // fn do_something2(&self, s: &str) -> &str {
    fn do_something2(&'b self, s: &str) -> &'b str {
        self.name
    }

    fn do_something3<'a>(&self, s: &'a str) -> &'a str {    // define lifetime 'a 
        s
    }
}

fn main() {
    let s = String::from("three");
    let a = StrA { name: &s };
    println!("{}", a.do_something());

    let s2 = String::from("hello");
    println!("{}", a.do_something2(&s2));
    println!("{}", a.do_something3(&s2));
}
```



## 6 - static lifetime

​	1、静态 lifetime

定义方式：	`'static`

其 lifetime 存活于整个执行期间， 所有的字符面值都拥有 `static` lifetime

`let s = &'static str = "hello";`

```rust
use std::fmt::Display;

fn function<'a, T: Display>(x: &'a str, y: &'a str, ann: T) -> &'a str {
    println!("ann = {}", ann);
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

fn function2<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

fn main() {
    let s1 = String::from("i am s1");
    let s2 = String::from("i am s2, hello");
    let ann = 129;
    let r = function(s1.as_str(), s2.as_str(), ann);
    println!("r = {}", r);

    let r2 = function2(s1.as_str(), s2.as_str());
    println!("r2 = {}", r2);
}
```







