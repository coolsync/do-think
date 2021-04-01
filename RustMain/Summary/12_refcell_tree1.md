# day20



## 1 [`RefCell` and the Interior Mutability Pattern](https://doc.rust-lang.org/book/ch15-05-interior-mutability.html#refcellt-and-the-interior-mutability-pattern)





1、内部可变性：允许在使用不可变引用时改变数据。



2、通过RefCell<T>在运行时检查借用规则（通常情况下，是在编译时检查借用规则），RefCell<T>代表其数据的唯一所有权。

类似于Rc<T>，RefCell<T>只能用于单线程场景。



3、选择Box<T>、Rc<T>或RefCell<T>的理由：

Rc<T> 允许相同数据有多个所有者；Box<T> 和 RefCell<T> 有单一所有者。

Box<T> 允许在编译时执行不可变或可变借用检查；Rc<T>仅允许在编译时执行不可变借用检查；RefCell<T> 允许在运行时执行不可变或可变借用检查。

因为 RefCell<T> 允许在运行时执行可变借用检查，所以我们可以在即便 RefCell<T> 自身是不可变的情况下修改其内部的值。



```rust
use std::rc::Rc;
use std::cell::RefCell;

#[derive(Debug)]
enum List {
    Cons(Rc<RefCell<i32>>, Rc<List>),
    Nil,
}

use crate::List::{Cons, Nil};

fn main() {
    let value = Rc::new(RefCell::new(5));

    let a = Rc::new(Cons(Rc::clone(&value), Rc::new(Nil)));
    let b = Cons(Rc::new(RefCell::new(6)), Rc::clone(&a));
    let c = Cons(Rc::new(RefCell::new(7)), Rc::clone(&a));

    println!("a after {:?}", a);
    println!("b after {:?}", b);
    println!("c after {:?}", c);
    println!("++++++++++++++++++++++");

    *value.borrow_mut() += 10;
    println!("a after {:?}", a);
    println!("b after {:?}", b);
    println!("c after {:?}", c);
}
```



## 2 [Creating a Reference Cycle](https://doc.rust-lang.org/book/ch15-06-reference-cycles.html#creating-a-reference-cycle)



![Reference cycle of lists](https://doc.rust-lang.org/book/img/trpl15-04.svg)

Figure 15-4: A reference cycle of lists `a` and `b` pointing to each other



```rust
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
```



# day21 



## 3 [Preventing Reference Cycles: Turning an `Rc` into a `Weak`](https://doc.rust-lang.org/book/ch15-06-reference-cycles.html#preventing-reference-cycles-turning-an-rct-into-a-weakt)



```rust
//弱引用Weak<T>
//特点：
//（1）弱引用通过Rc::downgrade传递Rc实例的引用，调用Rc::downgrade会得到Weak<T>类型的智能指针，同时将weak_count加1（不是将strong_count加1）。
//（2）区别在于 weak_count 无需计数为 0 就能使 Rc 实例被清理。只要strong_count为0就可以了。
//（3）可以通过Rc::upgrade方法返回Option<Rc<T>>对象。

use std::cell::RefCell;
use std::rc::Rc;
use std::rc::Weak;

#[derive(Debug)]
enum List {
    Cons(i32, RefCell<Weak<List>>),
    Nil,
}

use crate::List::{Cons, Nil};

impl List {
    fn tail(&self) -> Option<&RefCell<Weak<List>>> {
        match self {
            Cons(_, item) => Some(item),
            Nil => None,
        }
    }
}
fn main() {
    let a = Rc::new(Cons(5, RefCell::new(Weak::new())));
    println!("1, a strong count: {:?}, a weak count: {:?}", Rc::strong_count(&a), Rc::weak_count(&a));
    println!("1, a tail: {:?}", a.tail());
    
    println!("+++++++++++++++++++++++++");

    let b = Rc::new(Cons(10, RefCell::new(Weak::new())));
    if let Some(link) = b.tail() {
        *link.borrow_mut() = Rc::downgrade(&a); // b->a
    }
    println!("2, a strong count: {:?}, a weak count: {:?}", Rc::strong_count(&a), Rc::weak_count(&a));
    println!("2, b strong count: {:?}, b weak count: {:?}", Rc::strong_count(&b), Rc::weak_count(&b));
    println!("2, b tail: {:?}", b.tail());
    
    println!("+++++++++++++++++++++++++");

    if let Some(link) = a.tail() {
        *link.borrow_mut() = Rc::downgrade(&b); // a->b
    }
    println!("3, a strong count: {:?}, a weak count: {:?}", Rc::strong_count(&a), Rc::weak_count(&a));
    println!("3, b strong count: {:?}, b weak count: {:?}", Rc::strong_count(&b), Rc::weak_count(&b));
    
    println!("3, a tail: {:?}", a.tail());
}
```

​	

## 4 week tree 1



```rust

use std::rc::{Rc, Weak};
use std::cell::RefCell;

#[derive(Debug)]
struct Node {
    value: i32,
    parent: RefCell<Weak<Node>>,
    children: RefCell<Vec<Rc<Node>>>
}

fn main() {
   let leaf = Rc::new(Node{
        value: 5,
        parent: RefCell::new(Weak::new()),
        children: RefCell::new(vec![]),   
   });

   println!("leaf parent {:#?}", leaf.parent.borrow().upgrade());

   let branch = Rc::new(Node{
       value: 3,
       parent: RefCell::new(Weak::new()),
       children: RefCell::new(vec![Rc::clone(&leaf)]),
   });
   *leaf.parent.borrow_mut() = Rc::downgrade(&branch);

   println!("leaf parent: {:#?}", leaf.parent.borrow().upgrade());
}
```



result:

```bash
leaf parent None
leaf parent: Some(
    Node {
        value: 3,
        parent: RefCell {
            value: (Weak),
        },
        children: RefCell {
            value: [
                Node {
                    value: 5,
                    parent: RefCell {
                        value: (Weak),
                    },
                    children: RefCell {
                        value: [],
                    },
                },
            ],
        },
    },
)
```



