#[derive(Debug)]

struct Dog {
    name: String,
    age: u8,
    weight: f32,
}

// method
impl Dog {
    fn get_name(&self) -> &str {
        // &self.name[..]   // ?
        &(self.name[..])
    }
    fn get_age(&self) -> u8 {
        self.age
    }

    // fn get_weight(&self) -> f32 {
    //     self.weight
    // }

    fn show(&self) {
        println!("ho ho ho");
    }
}

// 分开method
impl Dog {
    fn get_weight(&self) -> f32 {
        self.weight
    }
}

fn main() {
    let dog = Dog {
        name: String::from("pangzhi"),
        age: 1,
        weight: 30.8,
    };

    println!("dog = {:#?}", dog);

    // Dog::show();
    dog.show();
    println!("name = {}", dog.get_name());
    println!("age = {}", dog.get_age());
    println!("weight = {}", dog.get_weight());
}
