# day31



## hfn1	[Function Pointers](https://doc.rust-lang.org/book/ch19-05-advanced-functions-and-closures.html#function-pointers)



```rust
//1、func 指向器

//func 指向器 允许我们使用 function 作为另一个 function 的参数。

// function 的类型是 fn ，fn 被称为 func 指向器 。指定参数为func 指向器 的语法类似于闭包。

fn add_one(val: i32) -> i32 {
    val + 1
}

fn do_twice(f: fn(i32) -> i32, val: i32) -> i32 {
    f(val) + f(val)
}

fn wrapper<T>(t: T, v: i32) -> i32
where
    T: Fn(i32) -> i32,      // 指定 generic
{
    t(v)
}

fn fn1(v: i32) -> i32 {
    v + 1
}

fn main() {
    // 传入 func pointer
    let r = do_twice(add_one, 5);
    println!("r: {}", r); // 12

    // +++++++++++++
    // 传入 闭包
    let a = wrapper(|x| x + 1, 1);
    println!("a = {}", a);

    // 传入 func pointer
    let b = wrapper(fn1, 1);
    println!("b = {}", b);
    
    println!("Hello, world!");
}
```



补充：  Fn FnMut FnOnce



## hfn2	[Returning Closures](https://doc.rust-lang.org/book/ch19-05-advanced-functions-and-closures.html#returning-closures)



```rust
// 2 return closures

// fn return_clo() -> Fn(i32) -> i32 { // trait objects without an explicit `dyn` are deprecated
//     |x| x+1
// }

fn return_clo() -> Box<dyn Fn(i32) -> i32> {
    Box::new(|x| x+1)
}

fn main() {
    let c = return_clo();   // return 闭包
    println!("1 + 1 = {}", c(1));
    println!("1 + 1 = {}", (*c)(1));

    println!("Hello, world!");
}
```

