use std::fs;
use std::fs::File;
use std::io;
use std::io::Read;
fn main() { 
    let r = read_username_from_file4();
    // match r {
    //     Ok(s) => println!("read file: {}", s),
    //     Err(e) => panic!("read file err: {:?}", e)
    // };
    println!("r = {:?}", r);
    println!("Hello, world!");
}

fn read_username_from_file() -> Result<String, io::Error> {
    let f = File::open("h1.txt");

    let mut f = match f {
        Ok(file) => file,
        Err(e) => return Err(e),
    };

    let mut s = String::new();

    match f.read_to_string(&mut s) {
        Ok(_) => Ok(s),
        Err(e) => Err(e),
    }
}

// A Shortcut for Propagating Errors: the ? Operator
fn read_username_from_file2() -> Result<String, io::Error> {
    let mut f = File::open("h1.txt")?;

    let mut s = String::new();

    f.read_to_string(&mut s)?;

    Ok(s)
}

// 在 ? 之后直接使用链式方法调用来进一步缩短代码
fn read_username_from_file3() -> Result<String, io::Error> {
    let mut s = String::new();

    File::open("h1.txt")?.read_to_string(&mut s)?;

    Ok(s)
}

// 更短的写法
fn read_username_from_file4() -> Result<String, io::Error> {
    fs::read_to_string("h.txt") // io::Result<String>
}
// fs::read_to_string fn: 它会打开文件、新建一个 String、读取文件的内容，并将内容放入 String，接着返回它。