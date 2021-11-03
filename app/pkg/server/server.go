package server

import (
	"context"
	"net/http"
	"time"
)

const (
	defaultReadTimeout     = 10 * time.Second
	defaultWriteTimeout    = 10 * time.Second
	defaultAddr            = ":80"
	defaultShutdownTimeout = 5 * time.Second
)

type Server struct {
	s               *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// NewServer 初始化并启动路由
func NewServer(handler http.Handler, opts ...Option) *Server {
	httpServer := &http.Server{
		Addr:         defaultAddr,
		Handler:      handler,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
	}

	ser := &Server{
		s:               httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: defaultShutdownTimeout,
	}

	for _, opt := range opts {
		opt(ser)
	}

	go ser.start()
	return ser
}

// Start 调用时，使用 go 关键字
func (s *Server) start() {
	s.notify <- s.s.ListenAndServe()
	close(s.notify)
}

// Notify 通知关闭
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown 关闭 httpserver
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	return s.s.Shutdown(ctx)
}
