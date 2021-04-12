package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"net/http"
	"time"
	"together/app/internal/global"
	"together/app/internal/routers"
)

func main() {
	// 初始化配置文件
	initViper()
	r := routers.New()
	s := &http.Server{
		Addr:         global.AppConfig.Addr,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	s.ListenAndServe()
}

func initViper() {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("toml")
	v.AddConfigPath("../../configs")
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到
			panic("配置文件未找到")
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("读取配置文件失败: %s \n", err.Error()))
		}
	}

	// 监听文件变动
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name) // e.Name 是修改的文件
		fmt.Printf("%+v", e.Op)                     // e.Op 是动作
	})

	if err := v.Unmarshal(&global.AppConfig); err != nil {
		fmt.Println(err)
	}
	global.Viper = v

	global.Viper.Get("server.addr")
	fmt.Println("init viper........OK")
}
