package parser

import (
	"ciphermagic.cn/imooc-basic/crawler/config"
	"ciphermagic.cn/imooc-basic/crawler/engine"
	"regexp"
)

const cityListRe = `href="(.*www\.zhenai\.com/zhenghun/[0-9a-z]+)"[^>]*>([^<^征婚]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}
	return result
}
