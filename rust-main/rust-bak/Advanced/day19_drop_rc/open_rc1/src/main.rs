// 通过Rc<T> option programs 的多个部分之间只读的共享data，因为相同位置的多个可变引用可能会造成data竞争和不一致。
use std::rc::Rc;

enum List {
    Cons(i32, Rc<List>),
    Nil,
}

use crate::List::{Cons, Nil};

fn main() {
    let a = Rc::new(Cons(5, Rc::new(Cons(10, Rc::new(Nil)))));
    println!("count after creating a, a count {}", Rc::strong_count(&a));

    let b = Cons(3, Rc::clone(&a));
    println!("count after bind to b, a count {}", Rc::strong_count(&a));

    {
        let c = Cons(4, Rc::clone(&a));
        println!("count after bind to c, a count {}", Rc::strong_count(&a));
    } // at the position, leave scope, c drop, counter - 1
    
    println!("count at end, a count {}", Rc::strong_count(&a));

    println!("Hello, world!");
}
