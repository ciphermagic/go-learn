package main

import (
	"ciphermagic.cn/imoocbasic/crawler/config"
	"ciphermagic.cn/imoocbasic/crawler/engine"
	"ciphermagic.cn/imoocbasic/crawler/persist"
	"ciphermagic.cn/imoocbasic/crawler/scheduler"
	"ciphermagic.cn/imoocbasic/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver(config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url:    "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}
