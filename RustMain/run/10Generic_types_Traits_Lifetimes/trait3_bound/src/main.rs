//1、trait_bound语法
//2、指定多个trait bound
//3、返回 trait的类型
// fn print_info(item: impl GetInfo) {  // work for params use
// fn print_info<T: GetInfo>(item: T) {  // use trait bound
//     println!("name: {}, age: {}", item.get_name(), item.get_age());
// }

trait GetName {
    fn get_name(&self) -> &String;
}

trait GetAge {
    fn get_age(&self) -> u32;
}

#[derive(Debug)]
pub struct Employee {
    pub name: String,
    pub age: u32,
}

impl GetName for Employee {
    fn get_name(&self) -> &String {
        &self.name
    }
}

impl GetAge for Employee {
    fn get_age(&self) -> u32 {
        self.age
    }
}

// write 1
// fn print_info<T: GetName + GetAge>(item: T) {
//     // use trait bound
//     println!("name: {}", item.get_name());
//     println!("age: {}", item.get_age());
// }

// write 2
fn print_info<T>(item: T)
where
    T: GetName + GetAge,
{
    println!("name: {}", item.get_name());
    println!("age: {}", item.get_age());
}

fn produce_item_with_age() -> impl GetAge {
    Employee{
        name: String::from("bob"),
        age: 30,
    }
}

fn main() {
    let emp = Employee {
        name: "bob".to_string(),
        age: 30,
    };
    print_info(emp);

    let e = produce_item_with_age();
    println!("{}", e.get_age());

    println!("Hello, world!");
}
