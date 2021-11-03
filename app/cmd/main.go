package main

import (
	"time"
	"together/app/internal/app"
	"together/configs"
	"together/log"

	"go.uber.org/zap"
)

// 仅执行读取配置文件
func main() {
	var cfg configs.AppServer
	setupConfig(&cfg)

	log.New("./logs/")
	defer func() {
		_ = zap.S().Sync()
	}()
	app.Run(&cfg)
}

func setupConfig(cfg *configs.AppServer) {
	s := configs.LoadConfig("configs/", "../configs/", "../../configs/")
	err := s.Read("AppServer", cfg)
	if err != nil {
		panic(err)
	}
	cfg.ReadTimeout *= time.Second
	cfg.WriteTimeout *= time.Second
}
