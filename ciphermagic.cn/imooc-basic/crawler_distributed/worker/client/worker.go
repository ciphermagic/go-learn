package client

import (
	"ciphermagic.cn/imooc-basic/crawler/config"
	"ciphermagic.cn/imooc-basic/crawler/engine"
	"ciphermagic.cn/imooc-basic/crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(request engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(request)
		var sResult worker.ParseResult
		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}
