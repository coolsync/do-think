//RefCell\Rc\Box

//RefCell<T>/Rc<T> 与Mutex<T>/Arc<T>
//1、Mutex<T>提供内部可变性，类似于RefCell
//2、RefCell<T>/Rc<T>是非线程安全的， Mutex<T>/Arc<T>是线程安全的

// use std::rc::Rc;
use std::sync::Arc;
use std::sync::Mutex;
use std::thread;

fn main() {
    // let counter = Mutex::new(0);
    // let counter = Rc::new(Mutex::new(0)); // `Rc<Mutex<i32>>` threads is not safely
    let counter = Arc::new(Mutex::new(0));
    let mut handles = vec![];

    for i in 0..10 {
        let ct = Arc::clone(&counter);
        let handle = thread::spawn(move || {
            let mut num = ct.lock().unwrap();
            *num += 1;
            // println!("{}, arc: {:?}, {:?}", i, Arc::strong_count(&ct), Arc::weak_count(&ct));
        });

        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap(); // wait 对应 thread end
    }

    println!("result: {:?}", counter);
}

// fn main() {
//     // let counter = Mutex::new(0);
//     let counter = Rc::new(Mutex::new(0));  // `Rc<Mutex<i32>>` threads is not safely
//     let mut handles = vec![];

//     for _ in 0..10 {
//         let cnt = Rc::clone(&counter);
//         let handle = thread::spawn(move || {
//             let mut cnt = counter.lock().unwrap();
//             *cnt += 1;
//         });

//         handles.push(handle);
//     }

//     for handle in handles {
//         handle.join().unwrap(); // wait 对应 thread end
//     }

//     println!("result: {}", *counter.lock().unwrap());
// }

// fn main() {
//     let counter = Mutex::new(0);
//     let mut handles = vec![];

//     for _ in 0..10 {
//         let handle = thread::spawn(move || {
//             let mut cnt = counter.lock().unwrap();
//             *cnt += 1;
//         });

//         handles.push(handle);
//     }

//     for handle in handles {
//         handle.join().unwrap(); // wait 对应 thread end
//     }

//     println!("result: {}", *counter.lock().unwrap());
// }
