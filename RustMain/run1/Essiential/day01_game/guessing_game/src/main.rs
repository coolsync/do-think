
use std::io;
use rand::Rng;  // trit
use std::cmp::Ordering;

fn main() {  
    // gen scret number
    let scret_number = rand::thread_rng().gen_range(1..101);
    
    println!("scret number:{}", scret_number);

    loop {
        println!("Guessing a number.");
    
        let mut guess = String::new();
        
        io::stdin().read_line(&mut guess).expect("read line err");
        
        println!("you guess num: {}", guess);
        
        // trim 消除边缘多余的stuff
        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            // Err(_) => continue, // err no handle
            Err(_) => {
                println!("parse err");
                break
            },
        };
        
        match guess.cmp(&scret_number) {
            Ordering::Less => println!("too samll."),
            Ordering::Greater => println!("too big"),
            Ordering::Equal => {
                println!("you win");
                break
            }
        }
    }
}
