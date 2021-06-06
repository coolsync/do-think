//1、trait用于定义与其它类型共享的功能，类似于其它语言中的接口。
//（1）可以通过trait以抽象的方式定义共享的行为。
//（2）可以使用trait bounds指定泛型是任何拥有特定行为的类型。

// 2. define trait
pub trait GetInfo {
    fn get_name(&self) -> &String;
    fn get_age(&self) -> u32;
}

// 3. impl trait
pub struct Employee {
    pub name: String,
    pub age: u32,
}

impl GetInfo for Employee {
    fn get_name(&self) -> &String {
        &self.name
    }

    fn get_age(&self) -> u32 {
        self.age
    }
}

pub struct Manager {
    pub name: String,
    pub age: u32,
    pub department: String,
}

impl GetInfo for Manager {
    fn get_name(&self) -> &String {
        &self.name
    }

    fn get_age(&self) -> u32 {
        self.age
    }
}
// 4、默认实现：可以在定义trait的时候提供默认的行为，trait的类型可以使用默认的行为
// 5. trait work for params
fn print_info(item: impl GetInfo) {
    println!("name: {}, age: {}", item.get_name(), item.get_age());
}

fn main() {
    let emp = Employee{name: "bob".to_string(), age:30};
    let mgr = Manager{name:"mark".to_string(), age:32, department: "s1".to_string()};
    // println!("emp, name: {}, age: {}", emp.get_name(), emp.get_age());
    // println!("mgr, name: {}, age: {}", mgr.get_name(), mgr.get_age());
    print_info(emp);
    print_info(mgr);

    println!("Hello, world!");
}
