package configs

import "time"

type AppServer struct {
	Addr           string
	RunMode        string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	BlogServerAddr string
}

type BlogServer struct {
	Addr     string
	IxugoURL string
}

type ImServer struct {
	Addr string
}

type Log struct {
	Path string
}
