package service

import (
	"fmt"
	"together/configs"
	pb "together/proto"

	"google.golang.org/grpc"
)

type Service struct {
	cfg    *configs.AppServer
	client pb.BlogServerClient
}

func New(cfg *configs.AppServer) *Service {
	return &Service{
		cfg:    cfg,
		client: getGrpc(cfg.BlogServerAddr),
	}
}

func getGrpc(addr string) pb.BlogServerClient {
	fmt.Println("GRPC start")
	// 创建 gRPC Channel 与 gRPC Server 进行通信（需服务器地址和端口作为参数）
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	return pb.NewBlogServerClient(conn)
}
