use std::sync::mpsc;
use std::thread;

fn main() {
    // create channel
    let (tx, rx) = mpsc::channel();

    thread::spawn(move || {
        let val = String::from("hi");
        tx.send(val).unwrap();
        // println!("val: {}", val); // 调用send的时候，会发生move动作，所以此处不能再使用val
    });

    let re = rx.recv().unwrap();
    println!("Recv Get: {}", re);
}
