// fn largest<T: PartialOrd+Copy>(list: &[T]) -> T {

fn largest<T>(list: &[T]) -> T      // 推荐
where
    T: PartialOrd + Copy,
{
    let mut larger = list[0];
    for &item in list.iter() {
        if item > larger {
            larger = item;
        }
    }
    larger
}

fn main() {
    let number_list = vec![10, 22, 8, 11];
    let max_number = largest(&number_list);
    println!("max num = {}", max_number);

    let char_list = vec!['a', 'y', 'b'];
    let max_char = largest(&char_list);
    println!("max char = {}", max_char);

    println!("Hello, world!");
}
