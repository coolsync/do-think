# day17

# Smart Pointer



## 1 Simple Introduce



1、指针是一个包含内存地址的变量。这个地址指向一些其它的数据。
 智能指针是一类数据结构，它们表现类似于指针，但是也拥有额外的元数据，最明显的，它们拥有一个引用计数。引用计数记录智能指针总共有多少个所有者，并且当没有任何所有者时清除数据。
 普通引用和智能指针的一个额外区别是：引用只是只借用数据的指针，而智能指针则是拥有它们指向的数据。

2、智能指针通常使用结构体实现。智能指针区别于常规结构体的显著特征在于其实现了Deref和Drop trait。
 （1）Deref trait允许智能指针结构体实例表现的像引用一样，这样就可以编写即用于引用，又用于智能指针的代码。
 （2）Drop trait允许我们自定义当智能指针离开作用域时执行的代码。

3、几个标准库中的智能指针：
 Box<T>，用于在堆上分配；
 Rc<T>，一个引用计数类型，其数据可以有多个所有者；
 Ref<T>和RefMut<T>，通过RefCell<T访问>，一个在运行时而不是在编译时执行借用规则的类型


链接：https://www.jianshu.com/p/8f41a4a4f9db





# day18 



## 1 Use Box

```rust
// 1| 最简单最直接的智能指针是box，其类型为Box<T>。box允许将值放在堆上而不是栈上，留着栈上的则是指向堆数据的指针。除了数据被存储在堆上外，box没有任何性能损失。
//
// 2| box适合用于如下场景：
// (1) 当有一个在编译时未知大小的类型，而又需要再确切大小的上下文中使用这个类型值的时候；（举例子：在一个list环境下，存放数据，但是每个元素的大小在编译时又不确定）
// (2) 当有大量数据并希望在确保数据不被拷贝的情况下转移所有权的时候；
// (3) 当希望拥有一个值并只关心它的类型是否实现了特定trait而不是其具体类型时。

// 3| use box:

fn main() {
    let b = Box::new(5);    // b store at stack, 5 store at heap, b point to 5 memory addr 
    println!("b = {}", b);
}
```



## 2 Box List size

```rust
//box适合用于如下场景：

// (1)当有一个在编译时未知大小的类型，而又需要再确切大小的上下文中使用这个类型值的时候；（举例子：在一个list环境下，存放数据，但是每个元素的大小在编译时又不确定）
// (2)当有大量数据并希望在确保数据不被拷贝的情况下转移所有权的时候；
// (3)当希望拥有一个值并只关心它的类型是否实现了特定trait而不是其具体类型时。

// enum List {
//     Cons(i32, List),
//     Nil,
// }
/*
 --> src/main.rs:7:1
  |
7 | enum List {
  | ^^^^^^^^^ recursive type has infinite size
8 |     Cons(i32, List),
  |               ---- recursive without indirection
  |
help: insert some indirection (e.g., a `Box`, `Rc`, or `&`) to make `List` representable
  |
8 |     Cons(i32, Box<List>),
  |               ^^^^    ^

*/

enum List {
    Cons(i32, Box<List>),
    Nil,
}

fn main() {
    use List::Cons;
    use List::Nil;
    // let l = Cons(1, Cons(2, Cons(3, Nil)));

    let l = Cons(1, 
        Box::new(Cons(2,
            Box::new(Cons(3, 
                Box::new(Nil))))));
}
```



## 3 Defer trait



```rust
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
```



## 4 Costum Box

```rust
// Costum Box 
use std::ops::Deref;

struct MyBox<T>(T);

impl<T> MyBox<T> {
    fn new(x: T) -> MyBox<T> {
        MyBox(x)
    }
}

impl<T> Deref for MyBox<T> {
    type Target = T;
    fn deref(&self) -> &T {
        &self.0
    }
}

fn main() {
    let x = 5;
    let y = MyBox::new(x);

    assert_eq!(5, x);
    assert_eq!(5, *y);
}
```



## 5 解引用多态 与 可变性 交互

```rust
// Costum Box 
use std::ops::Deref;

struct MyBox<T>(T);

impl<T> MyBox<T> {
    fn new(x: T) -> MyBox<T> {
        MyBox(x)
    }
}

impl<T> Deref for MyBox<T> {
    type Target = T;
    fn deref(&self) -> &T {
        &self.0
    }
}

fn main() {
    let m = MyBox::new(String::from("Rust"));
    hello(&m); // 将 MyBox 变为 &String，再将 String 解引用，变为字符串 slice。  &str
}

fn hello(name: &str) {
    println!("Hello, {}", name);
}

// 解引用多态与可变性交互
// when T: Defer<Target:U>, from &T to &U,
// when T: Defer<Target:U>, from &mut T to &U, 
// when T: DeferMut<Target:U>, from &mut T to &mut U,
```

























