mod mod_a {
    #[derive(Debug)]
    pub struct A {
        pub num: i32,
        name: String,
    }

    impl A {
        pub fn new(num: i32) -> A {
            A {
                num: num,
                name: String::from("bob"),
            }
        }

        pub fn print_a(&self) {
            println!("num: {}, name:{}", self.num, self.name);
        }
    }

    pub mod mod_b {
        pub fn print_b() {
            println!("b");
        }
        pub mod mod_c {
            pub fn print_c() {
                println!("c");
                super::print_b();   // call father mod fn;
            }
        }
    }
}

use mod_a::A;
use mod_a::A as A1;

fn main() {
    // let mut a = mod_a::A::new(1);
    // let a = A::new(1);
    let a = A1::new(1);
    a.print_a();

    // a.num = 2;
    // a.name = String::from("mark");
    // a.print_a();

    println!("+++++++++++++++++++");
    mod_a::mod_b::mod_c::print_c();
    println!("Hello, world!");
}
