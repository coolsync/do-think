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

	// instance a grpc client
	cli := pb.NewUserInfoClient(conn)

	// 组装 request param
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// create context, pass into client method
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := cli.GetUserInfo(ctx, &pb.UserRequest{Name: name})
	if err != nil {
		log.Fatalf("cli side, get user info failed: %v", err)
	}

	// print resp msg
	log.Println("result: ", res)
}
