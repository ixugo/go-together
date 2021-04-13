package service

import (
	"context"
	"fmt"
	"net"
	pb "together/proto"

	"google.golang.org/grpc"
)

func New(addr string) {
	fmt.Println("run im server")

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterImServerServer(s, new(imServer))
	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}

type imServer struct{}

func (s *imServer) SayHello(ctx context.Context, r *pb.ImHelloRequest) (*pb.ImHelloReply, error) {
	return &pb.ImHelloReply{Message: "Hello " + r.GetName()}, nil
}
