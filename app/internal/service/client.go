package service

import (
	"context"
	pb "together/proto"
)

// SayHello 测试用例
func (s *Service) SayHello(ctx context.Context, name string) (*pb.HelloReply, error) {
	return s.client.SayHello(ctx, &pb.HelloRequest{Name: name})
}

// GetList 获取博客列表
func (s *Service) GetList(ctx context.Context, url string) (*pb.GetListReply, error) {
	return s.client.GetList(ctx, &pb.GetListRequest{Url: url})
}
