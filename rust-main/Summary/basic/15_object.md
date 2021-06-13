# day24



## Characteristics of Object-Oriented Languages



### oo1  simple implement

```rust
//对象、封装、继承

//对象：数据和操作数据的过程
//Rust里面，结构体、枚举类型加上impl块

struct Dog {
    name: String,
}

// method
impl Dog {
    fn print_name(&self) {
        println!("Dog name: {}", self.name);
    }
}

fn main() {
    let d = Dog{name: String::from("wangcai")};
    d.print_name();

    println!("Hello, world!");
}
```



### oo2 encapsulation

./Cargo.toml:

```toml
[workspace]
members = [
    "getaver",
    "main",
]
```



getaver/src/lib.rc

```rust
pub struct AverCollection {
    list: Vec<i32>,
    aver: f64,
}

impl AverCollection {
    pub fn new() -> AverCollection {
        AverCollection {
            list: vec![],
            aver: 0.0,
        }
    }

    pub fn add(&mut self, value: i32) {
        self.list.push(value);
        self.update_aver();
    }

    pub fn remove(&mut self) -> Option<i32> {
        let result = self.list.pop();
        match result {
            Some(value) => {
                self.update_aver();
                Some(value)
            },
            None => None,
        }
    }

    pub fn average(&self) -> f64 {
        self.aver
    }
    
    fn update_aver(&mut self) {
        let total: i32 = self.list.iter().sum();
        self.aver = total as f64 / self.list.len() as f64;
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}
```



main/Cargo.toml:

```toml
[package]
name = "main"
version = "0.1.0"
authors = ["dart"]
edition = "2018"

# See more keys and their definitions at https://doc.rust-lang.org/cargo/reference/manifest.html

[dependencies]
getaver = {path = "../getaver"}
```



main/src/main.rc:

```rust
use getaver;

fn main() {
    let mut a = getaver::AverCollection::new();

    a.add(1);
    println!("average: {}", a.average());
    
    a.add(2);
    println!("average: {}", a.average());

    a.add(3);
    println!("average: {}", a.average());

    a.remove(); // remove 3
    println!("average: {}", a.average());
}
```





### oo3 Inheritance

Rust里面没有继承的概念，可以通过tait来进行行为共享

```rust
// 伪代码

trait A {
	fn sum();
}

struct XXX {
}

impl A for XXX {
	fn sum() {
		//todo
	}
}
```

