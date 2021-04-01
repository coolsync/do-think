// 1 Option 是 std 内 一个 enum, like:
// enum Option<T> {
//     Some(T),
//     None
// }

// handler nil pointer.

// 2 usage 方式

fn main() {
    let some_num = Some(1);
    let some_str = Some(String::from("a string"));
    let absent_num: Option<i32> = None;

    let x: i32 = 5;
    // let y = Some(5);
    let y: Option<i32> = Some(5);

    // let sum = x + y;    // no implementation for `i32 + std::option::Option<{integer}>`

    let mut temp = 0;
    match y {
        Some(y) => {
            temp = y;
        }
        None => {
            println!("do nothing")
        }
    }

    // let result = plus_one(y);
    // match result {
    //     Some(i) => println!("reuslt: {}", i),
    //     None => println!("nothing"),
    // }
    
    if let Some(value) = plus_one(y) {
        println!("value: {}", value);
    } else {
        println!("do nothing");
    }

    let sum = x + temp;
    println!("sum = {}", sum);
}

fn plus_one(i: Option<i32>) -> Option<i32> {
    match i {
        Some(i) => Some(i + 1),
        None => None,
    }
}
