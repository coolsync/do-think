use std::thread;

fn main() {
    let v = vec![1, 2, 3];

    let handle = thread::spawn(move || {
        println!("v {:?}", v); // v move here
    }); // v drop

    handle.join().unwrap();
    // println!("v {:?}", v); // not print
    println!("Hello, world!");
}

// fn main() {
//     let v = vec![1, 2, 3];

//     let handle = thread::spawn(|| {
//         println!("v {:?}", v);
//     });
//     // drop(v)

//     handle.join().unwrap();
//     println!("Hello, world!");
// }
