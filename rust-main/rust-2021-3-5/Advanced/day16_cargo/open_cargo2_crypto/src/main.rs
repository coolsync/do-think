extern crate crypto;

use crypto::digest::Digest;
use crypto::sha3::Sha3;

fn main() {
    // 1. create hash obj
    let mut hasher = Sha3::sha3_256();
    // 2. input string to conv
    hasher.input_str("hello world!");
    // 3. get hash value
    let r = hasher.result_str();
    println!("r: {}", r);
}
