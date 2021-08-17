package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"together/app/internal/routers"
	rest "together/app/pkg/server"
	"together/configs"
	"together/global"
)

func main() {
	setupConfig()

	server := rest.NewServer(routers.New(),
		rest.Port(global.AppServer.Addr),
		rest.ReadTimeout(global.AppServer.ReadTimeout),
		rest.WriteTimeout(global.AppServer.WriteTimeout),
	)

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	select {
	case s := <-exit:
		fmt.Printf(fmt.Sprintf("s(%s) := <-interrupt ", s.String()))
	case err := <-server.Notify():
		fmt.Printf(fmt.Sprintf(` err(%s) = <-server.Notify()`, err))
	}

	if err := server.Shutdown(); err != nil {
		fmt.Printf(` err(%s) := server.Shutdown()`, err)
	}
}

func setupConfig() {
	s := configs.LoadConfig("configs/", "../configs/", "../../configs/")
	err := s.Read("AppServer", &global.AppServer)
	if err != nil {
		panic(err)
	}
	global.AppServer.ReadTimeout *= time.Second
	global.AppServer.WriteTimeout *= time.Second
}
