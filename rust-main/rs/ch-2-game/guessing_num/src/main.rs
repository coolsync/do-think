use std::io;

fn main() {
    println!("guessing number！");

    println!("guessing a number: ");

    let mut guess_num = String::new();

    io::stdin().read_line(&mut guess_num).expect("read line err");

    println!("you guessing number is: {}", guess_num)
}
