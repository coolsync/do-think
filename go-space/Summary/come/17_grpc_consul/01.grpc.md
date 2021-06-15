# GRPC



## 1 Create Proto File

path: day17/01grpc/proto/user.proto

```protobuf
syntax = "proto3";  // set version num

option go_package = "./user"; //  设置输出目录 

package user; // 设置包名

// Service Method
service UserInfo {
    // Method Name
    rpc GetUserInfo (UserRequest) returns (UserResponse);
}

// Define request struct
message UserRequest {
    string name = 1;
}

// Define response struct
message UserResponse {
    int32 id = 1;
    string name = 2;
    int32 age = 3;
    repeated string hobby = 4;
}

```



in day17/01grpc, run shell:

protoc --go_out=. --go-grpc_out=. proto/*.proto

generate:

user/user_grpc.pb.go

user/user.pb.go

## 2 server side main.go

```go
package main

import (
	"context"
	pb "day17/01grpc/user"
	"errors"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserInfoServer
}

const (
	port = ":50051"
)

// SayHello implements helloworld.GreeterServer
// func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
// 	log.Printf("Received: %v", in.GetName())
// 	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
// }
func (s *server) GetUserInfo(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	name := in.GetName()
	if name == "mark" {
		return &pb.UserResponse{
			Id:    1,
			Name:  name,
			Age:   30,
			Hobby: []string{"run", "dianyin"},
		}, nil
	}
	return nil, errors.New("no such user")
}

func main() {
	// 1 create listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("server side, failed to listen: %v", err)
	}
    
	// 2 instance grpc server
	s := grpc.NewServer()

	// 3 on grpc server, resgister remote call method
	pb.RegisterUserInfoServer(s, &server{})

	// 4 run service
	if err = s.Serve(lis); err != nil {
		log.Fatalf("server side, failed to serve: %v", err)
	}
}

```



## 3 client side main.go

```go
package main

import (
	"context"
	pb "day17/01grpc/user"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "mark"
)

func main() {
	// connect server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("cli side, dial failed: %v", err)
	}
	defer conn.Close()

	// instance a grpc client, new client in pb file
	cli := pb.NewUserInfoClient(conn)

	// 组装 request param
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// create context, pass to client method
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := cli.GetUserInfo(ctx, &pb.UserRequest{Name: name})
	if err != nil {
		log.Fatalf("cli side, get user info failed: %v", err)
	}

	// print resp msg
	log.Println("result: ", res)	
}
```





