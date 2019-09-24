package main

import (
	"ciphermagic.cn/imoocbasic/crawler/engine"
	"ciphermagic.cn/imoocbasic/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
