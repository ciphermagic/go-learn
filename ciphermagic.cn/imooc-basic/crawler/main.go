package main

import (
	"ciphermagic.cn/imooc-basic/crawler/config"
	"ciphermagic.cn/imooc-basic/crawler/engine"
	"ciphermagic.cn/imooc-basic/crawler/persist"
	"ciphermagic.cn/imooc-basic/crawler/scheduler"
	"ciphermagic.cn/imooc-basic/crawler/zhenai/parser"
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
