package service

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "together/proto"

	"google.golang.org/grpc"
)

func New(addr string) {
	fmt.Println("start")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBlogServerServer(s, new(blogServer))
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

type blogServer struct {
}

func (s *blogServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + r.GetName()}, nil
}
