//4、解引用raw指针
//不可变和可变的，分别写作*const T, *mut T
//
//(1)允许忽略借用规则，可以同时拥有不可变和可变的指针，或者是多个指向相同位置的可变指针
//（2）不保证指向的内存是有效的
//（3）允许为空
//（4）不能实现任何自动清理的功能
//

fn main() {
    let mut num = 5;

    let r1 = &num as *const i32;
    let r2 = &mut num as *mut i32;

    unsafe {
        println!("r1: {}", *r1);
        println!("r2: {}", *r2);
    }

    // err: dereference of raw pointer is unsafe and requires unsafe function or block
    
    // println!("r1: {}", *r1);
    // println!("r2: {}", *r2);

    let add = 0x12456789usize;
    let _r = add as *const i32;

    println!("Hello, world!");
}
