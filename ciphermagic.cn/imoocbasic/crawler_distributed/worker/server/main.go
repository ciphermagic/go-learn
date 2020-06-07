package main

import (
	"ciphermagic.cn/imoocbasic/crawler/config"
	"ciphermagic.cn/imoocbasic/crawler_distributed/rpcsupport"
	"ciphermagic.cn/imoocbasic/crawler_distributed/worker"
	"fmt"
	"log"
)

func main() {
	log.Fatal(serveRpc(fmt.Sprintf(":%d", config.WorkerPort0)))
}

func serveRpc(host string) error {
	return rpcsupport.ServeRpc(host, &worker.CrawlService{})
}
