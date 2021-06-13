// 对 任何 实现 特定 trait 类型，有条件实现另一个 trait method
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
