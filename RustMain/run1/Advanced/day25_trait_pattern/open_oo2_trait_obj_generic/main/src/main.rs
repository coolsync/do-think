use gui::{Screen, Button, SelectBox};

fn main() {
    let s = Screen{
        components: vec![
            Box::new(Button{
                width: 50,
                height: 10,
                label: String::from("Yes"),
            }),
            Box::new(SelectBox{
                width: 70,
                height: 20,
                option: vec![
                    String::from("Yes"),
                    String::from("No"),
                    String::from("MayBe"),
                ],
            }),
        ],
    };

    // Use generic // mismatched types expected struct `gui::Button`, found struct `gui::SelectBox`
    
    // let s = Screen{
    //     components: vec![
    //         Button{
    //             width: 50,
    //             height: 10,
    //             label: String::from("Yes"),
    //         },
    //         SelectBox{  
    //             width: 70,
    //             height: 20,
    //             option: vec![
    //                 String::from("Yes"),
    //                 String::from("No"),
    //                 String::from("MayBe"),
    //             ],
    //         },
    //     ],
    // };

    s.run();  
    println!("Hello, world!");
}
