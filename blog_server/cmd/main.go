package main

import (
	"github.com/elastic/go-elasticsearch/v7"
	"together/blog_server/internal/dao"
	"together/blog_server/internal/service"
	"together/configs"
	"together/global"
)

func main() {
	setupConfig()
	service.New(global.BlogServer.Addr)
}

func setupConfig() {
	s := configs.LoadConfig("configs/", "../configs/", "../../configs/")
	err := s.Read("BlogServer", &global.BlogServer)
	if err != nil {
		panic(err)
	}
	cfg := elasticsearch.Config{
		Addresses: global.BlogServer.EsAddresses,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	dao.EsClient = es
}
