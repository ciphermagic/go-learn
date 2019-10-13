package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)

	const resultSize = 494
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"阿坝",
		"阿克苏",
		"阿拉善盟",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("result shold hava %d requests; but had %d", resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; bug was %s", i, expectedUrls[i], result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result shold hava %d itmes; but had %d", resultSize, len(result.Items))
	}
	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("expected city #%d: %s; bug was %s", i, expectedCities[i], result.Items[i].(string))
		}
	}
}
