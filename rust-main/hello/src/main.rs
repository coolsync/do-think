use std::env;

fn main() {
   let name = env::args().skip(1).next();


   match name {     // options
       Some(n) => println!("hi {}", n),
       None => println!("None")
   }
}
