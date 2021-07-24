mod mod_a {
    #[derive(Debug)]    // 能够 print struct ino
    pub struct A {
        pub number: i32,
        name: String,
    }

    impl A {
        pub fn new_a() -> A {
            A {
                number: 1,
                name: String::from("A"),
            }
        }

        pub fn get_a(&self) {
            println!("number: {}, name: {}", self.number, self.name);
        }
    }

    pub mod mod_b {
        pub fn print_B() {
            println!("B");
        }

        pub mod mod_c {
            pub fn print_C() {
                println!("C");
                super::print_B();   // call Up-level function 
            }
        }
    }
}

// use mod_a::A;
use mod_a::A as A1;

fn main() {
    // let a = mod_a::A::new_a();    // 绝对 path
    // let a = A::new_a(); // 使用 use
    let a = A1::new_a();
    a.get_a();

    let number_a = a.number;
    // let name_a = a.name;

    println!("++++++++++++++");
    mod_a::mod_b::mod_c::print_C();
}
