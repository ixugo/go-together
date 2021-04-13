package main

import "together/im_server/internal/service"

func main() {
	service.New(":8082")
}
