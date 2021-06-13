mod factory {
    pub mod produce_refrigerator {  // 生产冰箱
        pub fn produce_re() {
            println!("produce refrigerator!");
        }
    }

    mod produce_washing_machines {  // 生产洗衣机   // 私有 不可调用
        fn produce_washing() {
            println!("produce washing machine!");
        }
    }
}

fn main() {
    factory::produce_refrigerator::produce_re();

    println!("Hello, world!");
}
