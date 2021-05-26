//5、调用 no safe 的 function 或者 method
unsafe fn dangerous() {
    println!("do something dangerous");
}

fn f1() {
    let mut num = 5;
    let r1 = &num as *const i32;
    let r2 = &mut num as *mut i32;

    unsafe {
        println!("*r1: {}", *r1);
        println!("*r2: {}", *r2);
    }
}
fn main() {
    unsafe {
        dangerous();
    }
    // dangerous(); // err

    f1();

    println!("Hello, world!");
}
