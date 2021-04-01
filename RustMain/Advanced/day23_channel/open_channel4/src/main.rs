use std::thread;
use std::sync::mpsc;
use std::time::Duration;

fn main() {
    // mutiple produce single custom
    let (tx, rx) = mpsc::channel();
    let tx1 = mpsc::Sender::clone(&tx);
    let tx2 = mpsc::Sender::clone(&tx);

    thread::spawn(move || {
        let vals = vec![
            String::from("hi"),
            String::from("from"),
            String::from("the"),
            String::from("thread"),
        ];

        for val in vals {
            tx.send(val).unwrap();  // not use tx, main thread still wait
            thread::sleep(Duration::from_secs(1));
        }
    });

    thread::spawn(move || {
        let vals = vec![
            String::from("A"),
            String::from("B"),
            String::from("C"),
            String::from("D"),
        ];

        for val in vals {
            tx1.send(val).unwrap();
            thread::sleep(Duration::from_secs(1));
        }
    });

    thread::spawn(move || {
        let vals = vec![
            String::from("aa"),
            String::from("bb"),
            String::from("cc"),
            String::from("dd"),

        ];
        
        for val in vals {
            tx2.send(val).unwrap();
            // thread::sleep(Duration::from_secs(1));
        }
    });

    for recv in rx {
        println!("Recv Get: {}", recv);
    }

    println!("Hello, world!");
}
