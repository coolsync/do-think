pub mod animal;

#[cfg(test)]
mod tests {
    // use crate::animal::*;

    use crate::animal::dog;
    // use crate::animal::bird;

    
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
    #[test]
    fn use_dog() {
        assert_eq!(true, dog::is_dog())
    }

    #[test]
    fn use_bird() {
        assert_eq!(true, crate::animal::bird::is_bird())
    }
}
