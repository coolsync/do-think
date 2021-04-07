#[macro_export]
macro_rules! my_vec {
    ($($x: expr), *) => {
        {
            let mut tmp_vec = Vec::new();
            $(
                tmp_vec.push($x);
            )*
            tmp_vec
        }
    };
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}
