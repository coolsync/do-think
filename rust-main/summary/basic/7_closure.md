# day14 Closures



path: `/home/dart/DoThinking/RustMain/Advanced/day14_closures/open_closure`

```rust
// 1| 闭包 define: 闭包是可以保存进变量或者作为参数传递给其它函数的匿名函数。
//    闭包和函数不同的是，闭包允许捕获调用者作用域中的值。

// 2| use 闭包
// 3| use 带有 generic and fn trait 的 闭包

// Statement format
fn add_one_v1(x: u32) -> u32 {
    x
}

fn main() {
    let use_closure = || println!("this closure");
    use_closure();

    let add_one_v2 = |x: u32| -> u32 {x+1};
    let add_one_v3 = |x| {x+1};
    let add_one_v4 = |x| x+1;
    
    let a = add_one_v1(5);
    let b = add_one_v2(5);
    let c = add_one_v3(5);
    let d = add_one_v4(5);
    println!("a = {}, b = {}, c = {}, d = {}", a, b, c, d);

    // 类型 不能 推导 两次的 example
    let example_closure = |x| x;
    
    let s = example_closure(String::from("hello"));
    println!("s: {}", s);
    
    // let n = example_closure(5);
    let n = example_closure(5.to_string());
    println!("n: {}", n);

    // 捕捉 环境变量的 值
    let i = 5;
    let get_env_val = |x| x+i;
    let r = get_env_val(1);
    println!("r: {}", r);
}
```



path: `/home/dart/DoThinking/RustMain/Advanced/day14_closures/open_closure_cacher`



```rust
// 实现 一个 缓冲， 处理并保存 第一次输入的值
struct Cacher<T>
where
    T: Fn(u32) -> u32,
{
    calculation: T,
    value: Option<u32>,
}

impl<T> Cacher<T>
where
    T: Fn(u32) -> u32,
{
    fn new(calculation: T) -> Cacher<T> {
        Cacher {
            calculation,
            value: None,
        }
    }

    fn value_cacher(&mut self, arg: u32) -> u32 {
        match self.value {
            Some(v) => v,
            None => {
                let v = (self.calculation)(arg);
                self.value = Some(v);
                v
            }
        }
    }
}

fn main() {
    let mut c = Cacher::new(|x| x+1);
    
    let r = c.value_cacher(5);
    println!("r: {}", r);

    let r2 = c.value_cacher(6);
    println!("r2: {}", r2);
}
```



path:  `/home/dart/DoThinking/RustMain/Advanced/day14_closures/open_closure_impl`



```rust
fn main() {
    // let x = 1;
    // let eqult_to_x = |z| x==z;
    // let y = 1;
    // assert!(eqult_to_x(y));

    let x = vec![1,2,3];
    let eqult_to_x = move |z| x==z; // move after, x drop
    // println!("x: {:?}", x); 

    let y = vec![1,2,3];
    assert!(eqult_to_x(y));
}
//闭包可以通过三种方式捕获其环境，它们对应 function 的三种获取参数的方式， 分别是获取所有权、可变借用、不可变借用。

//这三种捕获值的方式被编码为如下三个Fn trait：
//（1）FnOnce消费从周围作用域捕获的变量，闭包周围的作用域被称为其环境。
// 为了消费捕获到的变量，闭包必须获取其所有权并在定义闭包时将其移进闭包。其名称的Once部分代表了闭包不能多次获取相同变量的所有权。
//（2）FnMut获取可变的借用值，所以可以改变其环境。
//（3）Fn从其环境获取不可变的借用值。
//当创建一个闭包时，rust会根据其如何使用环境中的变量来推断我们希望如何引用环境。由于所有闭包都可以被调用至少一次，因此所有闭包都实现了FnOnce。没有移动被捕获变量的所有权到闭包的闭包也实现了FnMut，而不需要对捕获的变量进行可变访问的闭包实现了Fn
```

