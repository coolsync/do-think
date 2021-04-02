# day25



## match pattern

```rust
// 1| 模式是Rust中特殊的语法，模式用来匹配值的结构。

// 2| A pattern consists of some combination of the following:

// Literals
// Destructured arrays, enums, structs, or tuples
// Variables
// Wildcards
// Placeholders

// match VALUE {
//     PATTERN => EXPRESSION,
//     PATTERN => EXPRESSION,
//     PATTERN => EXPRESSION,
// }

// pattern must match all codition
// fn main() {
//     let a = 1;
//     match a {
//         0 => println!("zero"),
//         1 => println!("one"),
//         _ => println!("other"),
//     }
//     println!("Hello, world!");
// }

// if let, has match, abort, back not run 
// fn main() {
//     let color: Option<&str> = None; // Mark

//     let is_ok = false;  
//     let age: Result<u8, _> = "33".parse();

//     if let Some(c) = color {
//         println!("color: {}", c);
//     } else if is_ok {
//         println!("is ok");
//     } else if let Ok(a) = age {
//         if a > 30 {
//             println!("mature man");
//         } else {
//             println!("young man");
//         }
//     } else {
//         println!("in else");
//     }
// }

// while let
//只要模式匹配就一直执行while循环
// fn main() {
//     let mut v = Vec::new();
//     v.push(1);
//     v.push(2);
//     v.push(3);

//     while let Some(top) = v.pop() {
//         println!("top: {}", top);
//     } // 只要匹配Some(value),就会一直循环
// }

// for
//在for循环中，模式是直接跟随for关键字的值，例如 for x in y，x就是对应的模式
// fn main() {
//     let v = vec!["a", "b", "c"];

//     for (index, value) in v.iter().enumerate() {    // enumerate get tuple
//         println!("index: {}, vlaue: {}", index, value);
//     }
// }
//此处的模式是(index, value)

// let
// let PATTERN = EXPRESSION
// fn main() {
// // (1, 2, 3) auto match (x, y, z), 1 bind to x, 2 bind to y,1 bind to z
//     let (x, y, z) = (1, 2, 3);
//     println!("{}, {}, {}", x, y, z);

//     let (x, .., z) = (1, 2, 3);
//     println!("{}, {}", x, z);
// }

// function
// function params 也是 模式
fn print_point(&(x, y): &(i32, i32)) {
    println!("x: {}, y: {}", x, y);
}

fn main() {
    let p = (1, 2);
    print_point(&p);
}

//模式在使用它的地方并不都是相同的，模式存在不可反驳的和可反驳的
```



# day26



## pattern 1 refutable（可反驳的）和 irrefutable（不可反驳的）



//1、模式有两种：refutable（可反驳的）和 irrefutable（不可反驳的）。能匹配任何传递的可能值的模式被称为是不可反驳的。对值进行匹配可能会失败的模式被称为可反驳的。
//
//2、只能接受不可反驳模式的有：函数、let语句、for循环。原因：因为通过不匹配的值程序无法进行有意义的工作。
//
//3、if let和while let表达式被限制为只能接受可反驳的模式，因为它们的定义就是为了处理有可能失败的条件。



```rust
fn main() {
    // let a: Option<i32> = Some(5); // match Some(value), None
    // let b: Option<i32> = None; // match Some(value), None

    // let Some(x) = a;    // refutable pattern in local binding: `None` not covered

    // if let Some(v) = a {
    if let v = 5 {      // warning: irrefutable to if-let pattern
        println!("v: {}", v);
    }
    println!("Hello, world!");
}

```



## pattern 2 literals, named variables, multiple pattern, ..

```rust
// 1| match literals (字面值)
// fn main() {
//     let x = 1;

//     match x {
//         1 => println!("one"),
//         2 => println!("two"),
//         _ => println!("other"),    
//     }
// }

// 2| match named variables (命名变量)
// fn main() {
//     let x = Some(5);
//     let y = 10; // position 1

//     match x {
//         Some(50) => println!("50"),
//         Some(y) => println!("y: {:?}", y), // 此处是 position 2
//         _ => println!("default case x = {:?}", x),
//     }

//     println!("x: {:?}, y: {:?}", x, y); // 此处是 position 1
// }

// 3| match 多个模式
// fn main() {
//     let x = 1;
//     match x {
//         1|2 => println!("1 or 2"),  // | 表示： match 1 or 2
//         3 => println!("3"),
//         _ => println!("other"),
//     };
// }

// 4| 通过 .. match
fn main() {
    // let x = 2;
    
    // match x {
    //     1..=5 => println!("1 to 5"),    // 1|2|3|4|5 => println!("1 to 5")
    //     _ => println!("ohter"),
    // };

    let x = 'c';
    
    match x {
        'a'..='j' => println!("1"),
        'k'..='z' => println!("2"),
        _ => println!("other"),
    }
}
```



## pattern 3

```rust
//解构并分解值
//解构元祖、结构体、枚举、引用
//
//解构结构体
struct Point {
    x: i32,
    y: i32,
}

fn main() {
    // let p = Point{x: 1, y: 2};
    // let Point{x: a, y: b} = p;  // variable a,b 分别匹配 x,y
    // assert_eq!(1, a);
    // assert_eq!(2, b);

    // let Point{x, y} = p; // let Point{x:x, y:y} = p;
    // assert_eq!(1, x);
    // assert_eq!(2, y);

    // 一部分 match
    let p = Point{x: 1, y: 0};
    match p {
        Point{x, y:0} => println!("x axis"),
        Point{x:0, y} => println!("y axis"),
        _ => println!("other"),
    };
    
    println!("Hello, world!");
}
```



## pattern 4 enum

```rust
enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
}
fn main() {
    let msg = Message::ChangeColor(0, 160, 255);

    match msg {
        Message::Quit => {
            println!("quit")
        },
        Message::Move { x, y } => println!("x: {}, y: {}", x, y),
        Message::Write(text) => println!("text: {}", text),
        Message::ChangeColor(r, g, b) => {
            println!("r: {}, g: {}, b: {}", r, g, b)
        },
    }

    println!("Hello, world!");
}

```



## pattern 5 [Destructuring Nested Structs and Enums](https://doc.rust-lang.org/book/ch18-03-pattern-syntax.html#destructuring-nested-structs-and-enums)



```rust
enum Color {
    Rgb(i32, i32, i32),
    Hsv(i32, i32, i32),
}

enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(Color),
}

fn main() {
    let msg = Message::ChangeColor(Color::Hsv(0, 160, 255));

    match msg {
        Message::ChangeColor(Color::Rgb(r, g, b)) => {
            println!("r: {}, g: {}, b: {}", r, g, b);
        },

        Message::ChangeColor(Color::Hsv(h, s, v)) => {
            println!("h: {}, s: {}, v: {}", h, s, v);
        },

        _ => ()
    }

    println!("Hello, world!");
}
```



## pattern 6 [Destructuring Structs and Tuples](https://doc.rust-lang.org/book/ch18-03-pattern-syntax.html#destructuring-structs-and-tuples)

```rust
struct Point {
    x: i32,
    y: i32,
}

fn main() {
    let ((a, b), Point{x, y}) = ((1,2), Point{x: 3, y: 4});

    println!("a: {}, b: {}, x: {}, y:{}", a, b, x, y);
    
    println!("Hello, world!");
}
```



## pattern 7 [Ignoring Values in a Pattern](https://doc.rust-lang.org/book/ch18-03-pattern-syntax.html#ignoring-values-in-a-pattern)



