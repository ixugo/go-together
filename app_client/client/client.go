package client

import (
	"log"
	pb "together/proto"

	"google.golang.org/grpc"
)

var Instance pb.BlogServerClient

func init() {
	conn, err := grpc.Dial(":6161", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	Instance = pb.NewBlogServerClient(conn)
}
