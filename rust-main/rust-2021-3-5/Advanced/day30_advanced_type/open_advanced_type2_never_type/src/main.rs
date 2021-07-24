//2、从不返回的never type
//Rust 有一个叫做 ! 的特殊类型。在类型理论术语中，它被称为 empty type，因为它没有值。
//我们更倾向于称之为 never type。在函数不返回的时候充当返回值

use rand::Rng;
use std::cmp::Ordering;
use std::io;

// 1. 生成随机数
// 2. 多次提示 user input
// 3. 创建 guess String 字符串，
// 保存用户输入内容
// 4. 转换 guess 成 u32, has err, 跳出此次 loop
// 5. 使用 cmp包 比较 guess 与 secret_num 是否一致
// 6, win, 使用 break 跳出 loop
fn main() {
    println!("Guess a num Game");

    let secret_num = rand::thread_rng().gen_range(1..101);

    println!("secret_num: {}", secret_num);

    loop {
        println!("Please a num: ");

        let mut guess = String::new();

        io::stdin()
            .read_line(&mut guess)
            .expect("read input content faild");

        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => continue, //continue 的值是 !。
                                //当 Rust 要计算 guess 的类型时，它查看这两个分支。
                                //前者是 u32 值，而后者是 ! 值。
                                //因为 ! 并没有一个值，Rust 决定 guess 的类型是 u32
        };

        match guess.cmp(&secret_num) {
            Ordering::Less => println!("too less"),
            Ordering::Greater => println!("too greater"),
            Ordering::Equal => {
                println!("you win");
                break;
            }
        }
    }
}


// // 例子2：panic!
// // Option 上的 unwrap 函数代码：

// impl<T> Option<T> { 
// 	pub fn unwrap(self) -> T { 
// 		match self { 
// 			Some(val) => val, 
// 			None => panic!("called `Option::unwrap()` on a `None` value"),
// 		} 
// 	}
// }

// // 说明：
// // match 时，Rust 知道 val 是 T 类型，panic! 是 ! 类型，所以整个 match 表达式的结果是 T 类型。