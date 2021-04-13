package configs

import "time"

type AppServer struct {
	Addr         string
	RunMode      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type BlogServer struct {
	Addr        string
	IxugoDomain string
}

type ImServer struct {
	Addr string
}

type Log struct {
	Path string
}
