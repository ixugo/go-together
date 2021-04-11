package main

import (
	"together/blog_server/internal/service"
)

const addr = ":8081"

func main() {
	service.New(addr)
}
