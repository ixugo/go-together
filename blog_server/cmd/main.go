package main

import (
	"together/blog_server/internal/service"
	"together/configs"
	"together/global"
)

func main() {
	setupConfig()
	service.New(global.BlogServer.Addr)
}

func setupConfig() {
	s := configs.LoadConfig("configs/", "../../configs/")
	err := s.Read("BlogServer", &global.BlogServer)
	if err != nil {
		panic(err)
	}
}
