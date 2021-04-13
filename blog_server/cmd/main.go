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
	s := configs.LoadConfig("configs/")
	err := s.Read("AppServer", &global.BlogServer)
	if err != nil {
		panic(err)
	}
}
