use add_one;
use add_two;

fn main() {
    let num = 10;
    println!("add_one: {} plus 1 eq {}", num, add_one::add_one(num));
    println!("add_two: {} plus 2 eq {}", num, add_two::add_two(num));
}
