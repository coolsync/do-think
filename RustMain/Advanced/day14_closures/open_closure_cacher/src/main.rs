// 实现 一个 缓冲， 处理并保存 第一次输入的值
struct Cacher<T>
where
    T: Fn(u32) -> u32,
{
    calculation: T,
    value: Option<u32>,
}

impl<T> Cacher<T>
where
    T: Fn(u32) -> u32,
{
    fn new(calculation: T) -> Cacher<T> {
        Cacher {
            calculation,
            value: None,
        }
    }

    fn value_cacher(&mut self, arg: u32) -> u32 {
        match self.value {
            Some(v) => v,
            None => {
                let v = (self.calculation)(arg);
                self.value = Some(v);
                v
            }
        }
    }
}

fn main() {
    let mut c = Cacher::new(|x| x+1);
    
    let r = c.value_cacher(5);
    println!("r: {}", r);

    let r2 = c.value_cacher(6);
    println!("r2: {}", r2);
}
