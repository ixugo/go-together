package main

import (
	"net/http"
	"time"
	"together/app/internal/routers"
)

const addr = ":8080"

func main() {
	r := routers.New()
	s := &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	s.ListenAndServe()
}
