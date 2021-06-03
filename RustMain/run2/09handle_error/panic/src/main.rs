use std::fs::File;

// 1. panic

// 2. Result 与可恢复的错误
// enum Result<T, E> {
//     Ok(T),
//     Err(E),
// }

fn main() {
    //
    // panic!("hello panic");
    // let v = vec![1,2,3];
    // v[99];

    // 2
    // let f:u32 = File::open("hello.txt");
    let f = File::open("hello.txt");

    // let f = match f {
    //     Ok(file) => file,
    //     Err(err) => {
    //         panic!("open file failed: {:?}", err);
    //     }
    // };

    let f = match f {
        Ok(file) => file,
        Err(error) => match error.kind() {
            // ErrorKind::NotFound => match File::create("hello.txt") {
            std::io::ErrorKind::NotFound => match File::create("hello.txt") {
                Ok(fc) => fc,
                Err(err) => panic!("create file failed {:?}", err),
            },
            other_error => panic!("open file failed {:?}", other_error),
        }
    };
    
    // println!("Hello, world!");
}
