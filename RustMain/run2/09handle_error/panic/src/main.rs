use std::fs::File;
use std::io::ErrorKind;
// 1. panic

// 2. Result 与可恢复的错误
// enum Result<T, E> {
//     Ok(T),
//     Err(E),
// }

fn main() {
    // 1
    // panic!("hello panic");
    // let v = vec![1,2,3];
    // v[99];

    // 2
    let f = File::open("hello.txt");

    // let f = match f {
    //     Ok(file) => file,
    //     Err(err) => panic!("open file err: {:?}", err),
    // };

    // let f = match f {
    //     Ok(file) => file,
    //     Err(error) => match error.kind() {
    //         ErrorKind::NotFound => match File::create("hello.txt") {
    //             Ok(fc) => fc,
    //             Err(e) => panic!("create file err: {:?}", e),
    //         },
    //         other_err => panic!("other error: {:?}", other_err),
    //     },
    // };

    
    // let f = File::open("h.txt").unwrap();
    let f = File::open("h.txt").expect("open h.txt failed");
    
    // 小结：
    // expect 与 unwrap 的使用方式一样：
    // 返回文件句柄或调用 panic! 宏。expect 用来调用 panic! 的错误信息将会作为参数传递给 expect，
    // 而不像unwrap 那样使用默认的 panic! 信息。

    
}
