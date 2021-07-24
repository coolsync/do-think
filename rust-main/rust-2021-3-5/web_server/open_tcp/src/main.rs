use std::net::{TcpListener, TcpStream};
use std::io::{Read, Write};
use std::fs;
fn handle_client(mut stream: TcpStream) {
    let mut buf = [0; 512];
    stream.read(&mut buf).unwrap();
    println!("req: {}", String::from_utf8_lossy(&buf[..]));

    let ct = fs::read_to_string("i.html").unwrap();
    let res = format!("HTTP/1.1 200 OK\r\n\r\n{}", ct);
    // let res = "HTTP/1.1 200 OK\r\n\r\n";
    stream.write(res.as_bytes()).unwrap();
    // stream.flush().unwrap();
}

fn main() -> std::io::Result<()> {
    let listener = TcpListener::bind("127.0.0.1:8081")?;

    // accept connections and process them serially
    for stream in listener.incoming() {
        handle_client(stream?);
    }
    Ok(())
}