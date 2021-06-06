// use mylib::factory::refrigerator_factory::produce;
// use mylib::factory::refrigerator_factory;
// use mylib::factory::washing_machine_factory;
use mylib::factory::*;
use mylib::factory::washing_machine_factory as W;
fn main() {
    // mylib::factory::refrigerator_factory::produce(); // absolute path
    // produce();
    refrigerator_factory::produce(); // relative path
    washing_machine_factory::produce();
    
    W::produce_washing_machines();
    println!("Hello, world!");
}
