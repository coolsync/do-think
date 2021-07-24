// 方法 lifetime
struct StrA<'a> {
    name: &'a str,
}

impl<'b> StrA<'b> {
    fn do_something(&self) -> u32 {
        3
    }

    // fn do_something2(&self, s: &str) -> &str {
    fn do_something2(&'b self, s: &str) -> &'b str {
        self.name
    }

    fn do_something3<'a>(&self, s: &'a str) -> &'a str {    // define lifetime 'a 
        s
    }
}

fn main() {
    let s = String::from("three");
    let a = StrA { name: &s };
    println!("{}", a.do_something());

    let s2 = String::from("hello");
    println!("{}", a.do_something2(&s2));
    println!("{}", a.do_something3(&s2));
}
