package parser

import (
	"ciphermagic.cn/imoocbasic/crawler/engine"
	"regexp"
)

const cityListRe = `href="(.*www\.zhenai\.com/zhenghun/[0-9a-z]+)"[^>]*>([^<^征婚]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
