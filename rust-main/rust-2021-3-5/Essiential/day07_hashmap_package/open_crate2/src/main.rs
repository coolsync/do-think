// use mylib::factory::produce_refrigerator;
// use mylib::factory::produce_refrigerator::produce_re;
// use mylib::factory::produce_washing_machines;
// use mylib::factory::produce_washing_machines as W;

use mylib::factory::*;

fn main() {
    mylib::factory::produce_refrigerator::produce_re(); // 绝对路径
    produce_refrigerator::produce_re(); // 使用 use， 推荐做法, path 到 调用的 上一级
    // produce_re();

    produce_washing_machines::produce_washing();
    produce_washing_machines::produce_re();
    
    // W::produce_washing();
    

    println!("Hello, world!");
}
