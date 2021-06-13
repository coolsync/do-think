mod factory {
    pub mod produce_refrigerator {
        pub fn produce_re() {
            println!("produce refrigerator");
        }
    }

    mod produce_washing_machines {
        fn produce_washing_machines() {
            println!("produce washing machines");
        }
    }
}

fn main() {
    factory::produce_refrigerator::produce_re();

    println!("Hello, world!");
}