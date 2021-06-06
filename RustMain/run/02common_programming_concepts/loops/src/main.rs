fn main() {
    // 1. Repeating Code with loop
    // loop {
    //     println!("again")
    // };

    // 2. Returning Values from Loops

    let mut count = 0;
    let result = loop {
        count += 1;
        if count == 10 {
            break count * 2;
        }
    };
    println!("result = {}", result);

    // Conditional Loops with while
    let a = [10, 20, 30, 40, 50];

    let mut index = 0;
    while index < 5 {
        println!("a[index] = {}", a[index]);
        index += 1

        // if index == 5 {
        //     break;
        // } // err
    }

    // Use for
    for elem in a.iter() {
        println!("elem: {}", elem)
    }

    // Range type
    for num in (1..4).rev() {
        println!("num: {}", num)
    }
    // let fib = ;
    // println!("fib = {}, {}, {}", fibonacci(8), fibonacci(3), fibonacci(2));
    println!("{} {} {} {} {}", fib(0), fib(1), fib(2), fib(3), fib(4));
    println!("Hello, world!");
}

// fn fibonacci(n: i32) -> i32 {
//     if n == 0 {
//         return 0;
//     } else if n == 1 || n == 2 {
//         return 1;
//     }
//     fibonacci(n - 1) + fibonacci(n - 2)
// }

fn fib(d: i32) -> i32 {
    if d == 0 || d == 1 {
        d
    } else {
        fib(d - 2) + fib(d - 1)
    }
}
