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

	// 3 on grpc server, resgister remote method
	pb.RegisterUserInfoServer(s, &server{})

	// 4 run service
	if err = s.Serve(lis); err != nil {
		log.Fatalf("server side, failed to serve: %v", err)
	}
}
