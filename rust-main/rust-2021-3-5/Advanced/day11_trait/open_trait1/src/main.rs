// Define trait
pub trait GetInfo {
    fn get_name(&self) -> &String;
    fn get_age(&self) -> u32;
}

// 默认实现
trait IncName {
    fn get_inc_name(&self) -> String {
        String::from("asia")
    }
}

// 实现 trait
pub struct Employee {
    name: String,
    age: u32,
}

pub struct Manager {
    name: String,
    age: u32,
    subject: String,
}

impl IncName for Employee {}

impl GetInfo for Employee {
    fn get_name(&self) -> &String {
        &self.name
    }

    fn get_age(&self) -> u32 {  // u32 可copy 类型
        self.age
    }
}

impl IncName for Manager {
    fn get_inc_name(&self) -> String {
        String::from("saner")
    }   
}

impl GetInfo for Manager {
    fn get_name(&self) -> &String {
        &self.name
    }
    fn get_age(&self) -> u32 {
        self.age
    }
}

// trait 作为 parameter
fn print_info(item: impl GetInfo) {
    println!("name: {}", item.get_name());
    println!("age: {}", item.get_age());
}

fn main() {
    let e = Employee{name: "jerry".to_string(), age: 18};
    let m = Manager{name: "alice".to_string(), age: 20, subject: String::from("natrual")};

    let employee_inc_name = e.get_inc_name();
    println!("employee inc name: {}", employee_inc_name);

    let manager_inc_name = m.get_inc_name();
    println!("manager inc name: {}", manager_inc_name);

    println!("Employee name : {}, age : {}", e.get_name(), e.get_age());
    println!("Manager name : {}, age : {}, subject : {}", m.get_name(), m.get_age(), m.subject);

    print_info(e);
    print_info(m);

    println!("Hello, world!");
}
