package main

import (
	server2 "ciphermagic.cn/go-micro-service/chatserver/server"
)

func main() {
	var s server2.Server
	s = server2.NewServer()
	s.Listen(":3333")
	s.Start()
}
