//! # My Crate
//! 
//! `my_crate` is a collection of utilities to make performing certain
//! calculations more convenient.

/// And one to the number given.
/// 
/// #Example
/// ```
/// let five = 5;
/// assert_eq!(6, my_crate::add_one(five));
/// ```

pub fn add_one(x: i32) -> i32 {
    x + 1
}

// #[cfg(test)]
// mod tests {
//     #[test]
//     fn it_works() {
//         assert_eq!(2 + 2, 4);
//     }
// }
