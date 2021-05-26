use std::net::{TcpListener, TcpStream};

fn handle_client(_stream: TcpStream) {
    // ...
    println!("receive some content ...")
}

fn main() -> std::io::Result<()> {
    let listener = TcpListener::bind("127.0.0.1:8081")?;
    // let listener = TcpListener::bind("127.0.0.1:80").unwrap();

    // accept connections and process them serially
    for stream in listener.incoming() {
        handle_client(stream?);
    }
    Ok(())
}
