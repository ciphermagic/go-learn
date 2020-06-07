package main

import (
	"ciphermagic.cn/imoocbasic/crawler/config"
	"ciphermagic.cn/imoocbasic/crawler_distributed/rpcsupport"
	"ciphermagic.cn/imoocbasic/crawler_distributed/worker"
	"fmt"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	var host = fmt.Sprintf(":%d", config.WorkerPort0)

	// start server
	go serveRpc(host)
	time.Sleep(time.Second)

	// start client
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// Call method
	req := worker.Request{
		Url: "http://localhost:8080/mock/album.zhenai.com/u/9040917498112347157",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "孤独比酒暖°柠小萌",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Errorf("result: %s; err: %s", result, err)
	} else {
		fmt.Println(result)
	}
}
