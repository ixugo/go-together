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
		// 创建 gRPC Channel 与 gRPC Server 进行通信（需服务器地址和端口作为参数）
		conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			panic(err)
		}
		blogServer = pb.NewBlogServerClient(conn)
	})
	return blogServer
}
