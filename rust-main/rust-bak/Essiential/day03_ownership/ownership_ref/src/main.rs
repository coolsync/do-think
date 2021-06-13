// fn main() {
//     let s1 = gives_onwership();
//     println!("s1 = {}", s1);

//     let mut s2 = String::from("hello s2");

//     let s3 = takes_and_give_back(s2);

//     s2 = takes_and_give_back(s3);
//     println!("s2 = {}", s2);    //  let mut s2, ok

//     // println!("s2 = {}", s2); // err, let s2, this value move, again ref it
// }

// fn gives_onwership() -> String {
//     let s = String::from("hello s1");
//     s
// }

// fn takes_and_give_back(s: String) -> String {
//     s
// }

// ref 引用
//引用: 用法&,
//让我们创建一个指向值的引用，但是并不拥有它，因为不拥有这个值，所以，当引用离开其值指向的作用域后也不会被丢弃

//借用:&mut

fn calcute_length(s: &String) -> usize {
    s.len()
}

fn modify_string(s: &mut String) {
    s.push_str(", world");
}

fn main() {
    let mut s1 = String::from("hello");
    // let s = &s1;
    // let len = calcute_length(s);
    // println!("len = {}", len);

    // modify_string(&mut s1);

    let r1 = &s1;
    let r2 = &s1;
    println!("{},{}", r1, r2);
    let r3 = &mut s1;

    println!("{}", r3);
    // println!("s1 = {}", s1);
}
