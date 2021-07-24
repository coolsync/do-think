// 1| HashMap<K, V>
// 2| 创建 HashMap
// 3| 读取
// 4| 遍历
// 5| 更新

use std::collections::HashMap;

fn main() {
    // 1 HashMap<K,V>
    
    // 2 create HashMap
    let mut scores: HashMap<String, i32> = HashMap::new();
    scores.insert(String::from("blue"), 10);
    scores.insert(String::from("red"), 20);

    let keys: Vec<String> = vec![String::from("blue"), String::from("red")];
    let values = vec![10, 20];

    let scores: HashMap<_,_> = keys.iter().zip(values.iter()).collect();

    // 3 read
    let k = String::from("blue");
    if let Some(v) = scores.get(&k) {   // get 返回的是一个 Option enum
        println!("v = {}", v);
    }


    let k = String::from("yellow");
    let v = scores.get(&k);
    match v {
        Some(value) => println!("value: {}", value),
        None => println!("None"),
    }

    // 4 traverse： 会以任意的顺序 遍历出来
    println!("+++++++++++");
    for (key, value) in &scores {
        println!("key: {}, value: {}", key, value);
    }
    println!("+++++++++++");

    // 5 update
    // 直接插入
    let mut ss: HashMap<String, i32>= HashMap::new();
    ss.insert(String::from("one"), 1);
    ss.insert(String::from("two"), 2);
    ss.insert(String::from("three"), 3);
    ss.insert(String::from("one"), 3);
    println!("ss = {:?}", ss);

    // 健不存在的时候才插入
    let mut ss1: HashMap<String, i32>= HashMap::new();
    ss1.insert(String::from("one"), 1);
    ss1.insert(String::from("two"), 2);
    ss1.insert(String::from("three"), 3);
    ss1.entry(String::from("one")).or_insert(3);
    println!("ss1 = {:?}", ss1);

    // due to old val update a new val
    let text = "hello world yes world";
    let mut map = HashMap::new();

    for word in text.split_whitespace() {
        let count = map.entry(word).or_insert(0);
        *count += 1;    // map{"hello": 1}
    }

    println!("map = {:?}", map);

    println!("Hello, world!");
}
