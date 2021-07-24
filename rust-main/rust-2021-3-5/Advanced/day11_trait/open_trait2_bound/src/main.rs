// 1| trait_bound 语法
// 2| 指定多个 trait bound
// 3| 返回 trait 的类型
// fn print_info(item: impl GetInfo) { // 直接作为 param 的写法
// fn print_info<T: GetInfo>(item: T) {    // 使用 trait bound 的写法
// println!("name: {}", item.get_name());
// println!("age: {}", item.get_age());
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

// 写法1
// fn print_info<T: GetName + GetAge>(item: T) {   // 使用 trait bound 的写法
//     println!("name: {}", item.get_name());
//     println!("age: {}", item.get_age());
// }

// 写法2
fn print_info<T>(item: T)
where
    T: GetName + GetAge,
{
    println!("name: {}", item.get_name());
    println!("age: {}", item.get_age());
}

// 作为返回值
fn produce_item_with_age() -> impl GetAge {
    Employee {
        name: String::from("bob"),
        age: 30,
    }
}

fn main() {
    let e = Employee {
        name: "jerry".to_string(),
        age: 18,
    };

    print_info(e);

    let e = produce_item_with_age();
    println!("Employee age = {}", e.get_age());

    println!("Hello, world!");
}
