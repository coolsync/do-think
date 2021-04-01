use std::rc::Rc;
use std::cell::RefCell; // runtime change immutable data

#[derive(Debug)]
enum List {
    Cons(i32, RefCell<Rc<List>>),
    Nil,
}

use crate::List::{Cons, Nil};

impl List {
    fn tail(&self) -> Option<&RefCell<Rc<List>>> {  // get tail list 
        match self {
            Cons(_, item) => Some(item),
            Nil => None,
        }
    }
}

fn main() {
    let mut a = Rc::new(Cons(5, RefCell::new(Rc::new(Nil))));
    println!("1, a rc count = {:?}", Rc::strong_count(&a));
    println!("1, a.tail = {:?}", a.tail());

    let b = Rc::new(Cons(10, RefCell::new(Rc::clone(&a)))); // b -> a
    println!("2, a rc count = {:?}", Rc::strong_count(&b));
    println!("2, b rc count = {:?}", Rc::strong_count(&b));
    println!("2, b.tail = {:?}", b.tail()); 

    if let Some(link) = a.tail() {
        *link.borrow_mut() = Rc::clone(&b); // a -> b
    }

    println!("3, a rc count = {:?}", Rc::strong_count(&a));
    println!("3, b rc count = {:?}", Rc::strong_count(&b));
    // println!("3, a.tail = {:?}", a.tail()); 
}
