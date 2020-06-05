package main

import (
	"ciphermagic.cn/imoocbasic/crawler/config"
	"ciphermagic.cn/imoocbasic/crawler/engine"
	"ciphermagic.cn/imoocbasic/crawler/model"
	"ciphermagic.cn/imoocbasic/crawler_distributed/rpcsupport"
	"fmt"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	var host = fmt.Sprintf(":%d", config.ItemSaverPort)

	// start ItemSaverServer
	go serveRpc(host, "test1")
	time.Sleep(time.Second)

	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// Call save
	item := engine.Item{
		Type: "zhenai",
		Url:  "http://localhost:8080/mock/album.zhenai.com/u/6294421564459125476",
		Id:   "6294421564459125476",
		Payload: model.Profile{
			Name:       "全球焦點猫儿.",
			Gender:     "女",
			Age:        34,
			Height:     126,
			Weight:     192,
			Income:     "2001-3000元",
			Marriage:   "离异",
			Education:  "博士及以上",
			Occupation: "财务",
			Hokou:      "苏州市",
			Xinzuo:     "双鱼座",
			House:      "租房",
			Car:        "有豪车",
		},
	}
	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
