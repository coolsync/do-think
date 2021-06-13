# day19

## Drop trait



### 1 drop 离开 scope 时, 清理代码



```rust
// 1| Drop trait like other lang 析构 function，
// 当值离开作用域的时候执行此func的代码。
struct Dog {
    name: String,
}

impl Drop for Dog {
    fn drop(&mut self) {
        println!("Dog {} leave", self.name);
    }
}

fn main() {
    let a = Dog{name: String::from("xiaopangzi")};
    {
        let b = Dog{name: String::from("dahuang")};
        println!("0 ++++++++++++++++");    
    }
    println!("1 ++++++++++++++++");

    println!("Hello, world!");
}
```



### 2 drop 提前释放



```rust
struct Dog {
    name: String,
}

impl Drop for Dog {
    fn drop(&mut self) {
        println!("{} leave.", self.name);
    }
}

// rust provide std::mem::drop(_x: T) method
fn main() {
    let a = Dog{name: String::from("wangcai")};
    let b = Dog{name: String::from("dahuang")};
    
    drop(b);
    drop(a);

    println!("+++++++++++++++++++");
    println!("Hello, world!");
}
```



## Rc

![Two lists that share ownership of a third list](https://doc.rust-lang.org/book/img/trpl15-03.svg)

Figure 15-3: Two lists, `b` and `c`, sharing ownership of a third list, `a`

### rc1

use Box

```rust
enum List {
    Cons(i32, Box<List>),
    Nil,
}

use crate::List::{Cons, Nil};

fn main() {
    let a = Cons(5, Box::new(Cons(10, Box::new(Nil))));
    let b = Cons(3, Box::new(a));
    let c = Cons(4, Box::new(a));
    println!("Hello, world!");
}
```

```bash
error[E0382]: use of moved value: `a`
  --> src/main.rs:11:30
   |
9  |     let a = Cons(5, Box::new(Cons(10, Box::new(Nil))));
   |         - move occurs because `a` has type `List`, which does not implement the `Copy` trait
10 |     let b = Cons(3, Box::new(a));
   |                              - value moved here
11 |     let c = Cons(4, Box::new(a));
   |                              ^ value used here after move

```



use Rc

```rust
use std::rc::Rc;
use crate::List::{Cons, Nil};
enum List {
    Cons(i32, Rc<List>),
    Nil,
}

fn main() {
    // let a = Cons(5, Box::new(Cons(10, Box::new(Nil))));
    let a = Rc::new(Cons(5, Rc::new(Cons(10, Rc::new(Nil)))));
    
    // let b = Cons(3, Box::new(a));
    // let b = Cons(3, Rc::clone(&a)); 
	let b = Cons(3, a.clone());    
    // let c = Cons(4, Box::new(a));
    let c = Cons(4, Rc::clone(&a));
    
    println!("Hello, world!");
}
```



### rc2

```rust
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
```

