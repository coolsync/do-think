# day15 Iterator

[**13.2.** Processing a Series of Items with Iterators](https://doc.rust-lang.org/nightly/book/ch13-02-iterators.html)

## Use



```rust
//1| 迭代器负责遍历序列中的每一项和决定序列何时结束的逻辑。
//2| 创建迭代器：迭代器是惰性的，意思就是在调用方法使用迭代器之前，不会有任何效果
//3| 每个迭代器都实现了iterator trait, iterator trait定义在标准库中：

//trait Iterator {
//    type Item;
//    fn next(mut self) -> Option<Self::Item>; //type Item和Self::Item这种用法叫做定义trait的关联类型
//}
//next是Iterator被要求实现的唯一的一个方法，next一次返回一个元素，当迭代器结束时候，返回None


fn main() {
    
    // 到目前为止， 不会 对 v1.iter 产生影响
    let v1 = vec![1,2,3];
    let mut v1_iter = v1.iter();

    // for v in v1_iter {
    //     println!("v = {}", v);
    // }
    
    if let Some(v) = v1_iter.next() {
        println!("v = {}", v);
    }
    if let Some(v) = v1_iter.next() {
        println!("v = {}", v);
    }
    if let Some(v) = v1_iter.next() {
        println!("v = {}", v);
    }
    if let Some(v) = v1_iter.next() {
        println!("v = {}", v);
    } else {
        println!("At end");
    }

    // -----迭代可变引用-----
    let mut v2 = vec![1,2,3];
    let mut v2_iter = v2.iter_mut();
    if let Some(v) = v2_iter.next() {
        *v = 5;
    }
    println!("v2 = {:?}", v2);

    // -----消费适配器-----
    let v1 = vec![1,2,3];
    let total: i32 = v1.iter().sum(); // 调用消费适配器 sum 求和
    println!("total: {}", total);    

    // -----迭代适配器-----
    println!("+++++++++++++++++++");
    let v1 = vec![1,2,3];
    println!("v1 = {:?}", v1);
    
    let v2: Vec<_> = v1.iter().map(|x| x+1).collect();
    println!("v2 = {:?}", v2);
    
    // ----- 过滤 -------
    println!("+++++++++++++++++++");
    let v1 = vec![1, 22, 3, 45];
    println!("v1 = {:?}", v1);
    
    let v2: Vec<_> = v1.into_iter().filter(|x| *x > 5).collect();
    println!("v2 = {:?}", v2);
}
```



## 自定义 迭代器

```rust

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
```

