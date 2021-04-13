package service

import (
	pb "together/proto"
)

// SayHello 测试用例
func (s *Service) SayHello(name string) (*pb.HelloReply, error) {
	return s.client.SayHello(s.ctx, &pb.HelloRequest{Name: name})
}

// GetList 获取博客列表
func (s *Service) GetList(url string) (*pb.GetListReply, error) {
	return s.client.GetList(s.ctx, &pb.GetListRequest{Url: url})
}
