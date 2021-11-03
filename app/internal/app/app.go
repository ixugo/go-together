package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"together/app/internal/http/routers"
	"together/app/pkg/server"
	"together/configs"
)

// Run 初始化依赖配置的相关组件及启动接口
func Run(cfg *configs.AppServer) {
	// 实例化数据库 orm
	// 传递给 service 层

	r := routers.New(cfg) // 将 service 对象传递过去
	ser := server.NewServer(r, server.Port(cfg.Addr))
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case s := <-exit:
		fmt.Printf("s(%s) := <-exit ", s.String())
	case err := <-ser.Notify():
		fmt.Printf(`err(%s) = <-server.Notify()`, err)
	}

	if err := ser.Shutdown(); err != nil {
		fmt.Printf(` err(%s) := server.Shutdown()`, err)
	}
}
