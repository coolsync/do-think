
struct Counter {
    count: u32,
}

impl Counter {
    fn new(count: u32) -> Counter {
        Counter {count}
    }
}

impl Iterator for Counter {
    type Item = u32;
    fn next(&mut self) -> Option<Self::Item> {
        self.count += 1;
        if self.count < 6 {
            Some(self.count)
        } else {
            None
        }
    }
}
fn main() {
    let mut counter = Counter::new(0);
    
    for i in 0..6 {     // 迭代 6 次
        if let Some(v) = counter.next() {
            println!("i: {}, v: {}", i, v);
        } else {
            println!("At end, i = {}", i);
        }
    }
}