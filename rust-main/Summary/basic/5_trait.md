

# day12

# Trait Codition Method

## 1. 使用 trait bound 有条件的实现方法

```rust
trait GetName {
    fn get_name(&self) -> &String;
}

trait GetAge {
    fn get_age(&self) -> u32;
}

struct HumanMatch<T, U> {
    master: T,
    emp: U,
}

// 聚合 两个 struct 的 trait的实现
impl<T: GetAge+GetName, U: GetAge+GetName> HumanMatch<T, U> {
    fn print_all_info(&self) {
        println!("master name: {}", self.master.get_name());
        println!("master age: {}", self.master.get_age());
        println!("emp name: {}", self.emp.get_name());
        println!("emp age: {}", self.emp.get_age());
    }
}

struct Master {
    name: String,
    age: u32,
}
impl GetName for Master {
    fn get_name(&self) -> &String {
        &self.name
    } 
}
impl GetAge for Master {
    fn get_age(&self) -> u32 {
        self.age
    }
}

struct Emp {
    name: String,
    age: u32,
}
impl GetName for Emp {
    fn get_name(&self) -> &String {
        &self.name
    } 
}
impl GetAge for Emp {
    fn get_age(&self) -> u32 {
        self.age
    }
}

fn main() {
    let m = Master{name: "bob".to_string(), age: 30};
    let e = Emp{name: "pual".to_string(), age: 20};
    let h = HumanMatch{master: m, emp: e};

    h.print_all_info();
    println!("Hello, world!");
}
```



## 2. 对 任何 实现 trait 类型，有条件实现另一个 trait method

```rust
trait GetName {
    fn get_name(&self) -> &String;
}

trait PrintName {
    fn print_name(&self);
}

impl<T: GetName> PrintName for T {
    fn print_name(&self) {
        println!("name = {}", self.get_name());
    }
}

struct People {
    name: String,
}

impl GetName for People {
    fn get_name(&self) -> &String {
        &self.name
    }
}

fn main() {
    let p = People{name: "mark".to_string()};
    p.print_name();
}
```





# Life circle

## 1 variable life circle



```rust
fn main() {
    // err example
    // let r;
    // {
    // let x = 5;
    // r = &x;
    // }
    // println!("r: {}", r);
    
    let r;

    let x = 5;
    r = &x;
    println!("r: {}", r);
}
```































