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
