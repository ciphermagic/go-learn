package main

import (
	"ciphermagic.cn/imoocbasic/crawler/engine"
	"ciphermagic.cn/imoocbasic/crawler/scheduler"
	"ciphermagic.cn/imoocbasic/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}

	e.Run(engine.Request{
		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
