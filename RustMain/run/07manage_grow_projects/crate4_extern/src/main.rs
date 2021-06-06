extern crate crypto;
// use crypto::digest::Digest;
use self::crypto::digest::Digest;
use self::crypto::sha3::Sha3;

fn main() {
    // create a SHA3-256 object
    let mut hasher = Sha3::sha3_256();

    // write input message
    hasher.input_str("abc");
    
    // read hash digest
    let hex = hasher.result_str();
    
    assert_eq!(hex, "3a985da74fe225b2045c172d6bd390bd855f086e3e9d525b46bfe24511431532");
    println!("hex: {}", hex);
    println!("Hello, world!");
}
