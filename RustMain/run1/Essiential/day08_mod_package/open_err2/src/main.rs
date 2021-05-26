use std::fs::File;
use std::io;
use std::io::Read;

fn main() {
    let r = read_username_from_file();
    match r {
        Ok(s) => println!("s = {}", s),
        Err(e) => println!("err = {:?}", e),
    }
}

// fn read_username_from_file() -> Result<String, io::Error> {
//     let f = File::open("h.txt");

//     let mut f = match f {
//         Ok(file) => file,
//         Err(e) => return Err(e),
//     };

//     let mut s = String::new();
//     match f.read_to_string(&mut s) {
//         Ok(_) => Ok(s),
//         Err(e) => Err(e),    
//     }
// }

// fn read_username_from_file() -> Result<String, io::Error> {
//     let mut f = File::open("h.txt")?;

//     let mut s = String::new();

//     f.read_to_string(&mut s)?;
//     Ok(s)
// }

fn read_username_from_file() -> Result<String, io::Error> {
    let mut s = String::new();

    File::open("h.txt")?.read_to_string(&mut s)?;
    Ok(s)
}