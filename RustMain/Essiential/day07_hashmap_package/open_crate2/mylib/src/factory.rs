pub mod produce_refrigerator {  // 生产冰箱
    pub fn produce_re() {
        println!("produce refrigerator!");
    }
}

pub mod produce_washing_machines {  // 生产洗衣机   // 私有 不可调用
    pub fn produce_washing() {
        println!("produce washing machine!");
    }

    pub fn produce_re() {
        println!("produce washing machine!");
    }
}