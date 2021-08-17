package rest

import (
	"context"
	"net/http"
	"time"
)

const (
	readTimeout     = 20 * time.Second
	writeTimeout    = 20 * time.Second
	shutdownTimeout = 3 * time.Second
	address         = ":1323"
)

// Server http 服务器
type Server struct {
	s               *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// NewServer .
func NewServer(handle http.Handler, opts ...Option) *Server {
	httpserver := &http.Server{
		Addr:         address,
		Handler:      handle,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	s := &Server{
		s:               httpserver,
		notify:          make(chan error),
		shutdownTimeout: shutdownTimeout,
	}

	for _, opt := range opts {
		opt(s)
	}

	go s.Start()
	return s
}

// Start 启动服务，调用时使用 go 关键字
func (s *Server) Start() {
	s.notify <- s.s.ListenAndServe()
	close(s.notify)
}

// Notify 关闭通知
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown 优雅的关闭 api
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	return s.s.Shutdown(ctx)
}
