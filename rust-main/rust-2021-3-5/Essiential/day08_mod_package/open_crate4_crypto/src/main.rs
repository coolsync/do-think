extern crate crypto;    // ref out pkg

use crypto::digest::Digest; // input_str
use crypto::sha3::Sha3;

fn main() {
    // get hash value
    let mut hasher = Sha3::sha3_512();
    hasher.input_str("hello world");
    let result = hasher.result_str();

    println!("hash: {}", result);
}