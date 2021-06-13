// implement Deref trait 允许我们重载解引用运算符。
//let a: A = A::new();//前提：A类型必须实现Deref trait
//let b = &a;
//let c = *b;//解引用

fn main() {
    let x = 5;  // stack, has Copy trait
    let y = &x;
    assert_eq!(5, x);
    assert_eq!(5, *y); // decode ref

    let z = Box::new(x); // copy x to heap
    assert_eq!(5, *z);  // decode ref
    
    println!("{:?}", z);    // 5
}
