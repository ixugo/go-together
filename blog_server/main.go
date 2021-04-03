package main

import (
	"context"
	"log"
	"net"
	pb "together/proto"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":6161")
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
