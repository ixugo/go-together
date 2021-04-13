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
	r := routers.New()
	s := &http.Server{
		Addr:         global.AppServer.Addr,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	s.ListenAndServe()
}

func setupConfig() error {
	s := configs.LoadConfig("configs/")
	err := s.Read("AppServer", &global.AppServer)
	if err != nil {
		panic(err)
	}
	global.AppServer.ReadTimeout *= time.Second
	global.AppServer.WriteTimeout *= time.Second
	return nil
}
