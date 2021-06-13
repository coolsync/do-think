
fn c_to_f(c :f32) -> f32 {
    c * 9.0 / 5.0  + 32.0
}

fn f_to_c(f :f32) -> f32 {
    (f-32.0) * 5.0 / 9.0 
}

fn main() {
    let c: f32 = 10.0;

    let f = c_to_f(c);

    println!("f: {}", f);

    let c2 = f_to_c(f);
    println!("c2: {}", c2);

}
