// 1. A User struct definition
struct User {
    username: String,
    email: String,
    sign_in_count: u64,
    active: bool,
}

fn main() {
    // 2. Creating an instance of the User struct
    let mut user1 = User {
        username: String::from("somename123"),
        email: String::from("some@example.com"),
        sign_in_count: 1,
        active: true,
    };

    user1.email = String::from("some1234@example.com");

    // 3. A build_user function that takes an email and username and returns a User instance
    let email = String::from("some@example.com");
    let usename = String::from("some123.com");

    let user2 = build_user(email, usename);

    // 5. 
    // let user2 = User {
    //     email: String::from("another@example.com"),
    //     username: String::from("anotherusername567"),
    //     active: user1.active,
    //     sign_in_count: user1.sign_in_count,
    // };
    let user2 = User {
        email: String::from("another@example.com"),
        username: String::from("anotherusername567"),
        ..user1
    };

    println!("Hello, world!");
}

// fn build_user(email: String, username: String) -> User {
//     User {
//         email: email,
//         username: username,
//         active: true,
//         sign_in_count: 1,
//     }
// }

// 4. Using the Field Init Shorthand when Variables and Fields Have the Same Name
fn build_user(email: String, username: String) -> User {
    User {
        email,
        username,
        active: true,
        sign_in_count: 1,
    }
}