package main

import (
	"ciphermagic.cn/imoocbasic/crawler_distributed/rpcsupport"
	"ciphermagic.cn/imoocbasic/crawler_distributed/worker"
	"flag"
	"fmt"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port)))
}

func serveRpc(host string) error {
	return rpcsupport.ServeRpc(host, &worker.CrawlService{})
}
