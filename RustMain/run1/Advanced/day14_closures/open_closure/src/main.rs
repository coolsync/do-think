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


