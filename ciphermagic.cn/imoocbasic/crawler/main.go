package main

import (
	"ciphermagic.cn/imoocbasic/crawler/engine"
	"ciphermagic.cn/imoocbasic/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
