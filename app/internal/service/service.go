package service

import (
	"context"
	"fmt"
	"sync"
	"together/global"
	pb "together/proto"

	"google.golang.org/grpc"
)

var (
	once       sync.Once
	blogServer pb.BlogServerClient
)

type Service struct {
	ctx    context.Context
	client pb.BlogServerClient
}

func New(ctx context.Context) Service {
	return Service{
		ctx:    ctx,
		client: getClient(global.AppServer.BlogServerAddr),
	}
}

func getClient(addr string) pb.BlogServerClient {
	once.Do(func() {
		fmt.Println("GRPC start")
		conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			panic(err)
		}
		blogServer = pb.NewBlogServerClient(conn)
	})
	return blogServer
}
