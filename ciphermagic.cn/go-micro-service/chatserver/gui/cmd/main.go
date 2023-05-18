package main

import (
	"ciphermagic.cn/go-micro-service/chatserver/client"
	"ciphermagic.cn/go-micro-service/chatserver/gui"
	"flag"
	"log"
)

func main() {
	address := flag.String("server", "127.0.0.1:3333", "address of server")
	flag.Parse()
	c := client.NewClient()
	err := c.Dial(*address)

	if err != nil {
		log.Fatal("Error when connect to server", err)
	}
	defer c.Close()

	go c.Start()
	gui.StartUi(c)
}
