package main

import (
	"net/http"
	"time"
	"together/app/internal/routers"
	"together/configs"
	"together/global"
)

func main() {
	setupConfig()
	s := http.Server{
		Addr:         global.AppServer.Addr,
		Handler:      routers.New(),
		ReadTimeout:  global.AppServer.ReadTimeout,
		WriteTimeout: global.AppServer.WriteTimeout,
	}
	s.ListenAndServe()
}

func setupConfig() {
	s := configs.LoadConfig("configs/", "../../configs/")
	err := s.Read("AppServer", &global.AppServer)
	if err != nil {
		panic(err)
	}
	global.AppServer.ReadTimeout *= time.Second
	global.AppServer.WriteTimeout *= time.Second
}
