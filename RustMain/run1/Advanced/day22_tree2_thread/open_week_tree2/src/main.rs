use std::rc::{Rc, Weak};
use std::cell::RefCell;

struct Node {
    value: i32,
    parent: RefCell<Weak<Node>>,
    children: RefCell<Vec<Rc<Node>>>,
}

fn main() {
    let leaf = Rc::new(Node{
        value: 3,
        parent: RefCell::new(Weak::new()),
        children: RefCell::new(vec![]),
    });
    println!(
        "1 leaf strong: {}, weak: {}",
        Rc::strong_count(&leaf), Rc::weak_count(&leaf)
    );
    {
        let branch = Rc::new(Node{
            value: 5,
            parent: RefCell::new(Weak::new()),
            children: RefCell::new(vec![Rc::clone(&leaf)]), // branch children -> leaf 
        });

        println!(
            "1 branch strong: {}, weak: {}",
            Rc::strong_count(&branch), Rc::weak_count(&branch)
        );

        *leaf.parent.borrow_mut() = Rc::downgrade(&branch); // leaf parent -> branch

        println!(
            "2 leaf strong: {}, weak: {}",
            Rc::strong_count(&leaf), Rc::weak_count(&leaf)
        );
        println!(
            "2 branch strong: {}, weak: {}",
            Rc::strong_count(&branch), Rc::weak_count(&branch)
        );
    }   // branch drop
    
    println!(
        "3 leaf strong: {}, weak: {}",
        Rc::strong_count(&leaf), Rc::weak_count(&leaf)
    );
    
    println!("Hello, world!"); 
}
