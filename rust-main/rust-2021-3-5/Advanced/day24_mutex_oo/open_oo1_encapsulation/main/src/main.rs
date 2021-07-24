use getaver;

fn main() {
    let mut a = getaver::AverCollection::new();

    a.add(1);
    println!("average: {}", a.average());
    
    a.add(2);
    println!("average: {}", a.average());

    a.add(3);
    println!("average: {}", a.average());

    a.remove(); // remove 3
    println!("average: {}", a.average());
}