package persist

import (
	"ciphermagic.cn/imooc-basic/crawler/engine"
	"ciphermagic.cn/imooc-basic/crawler/model"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"testing"
)

func TestSaver(t *testing.T) {
	expected := engine.Item{
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
	client, _ := elastic.NewClient(
		elastic.SetURL("http://192.168.199.137:9200"),
		elastic.SetSniff(false),
	)

	const index = "dating_test"
	// Save expected item
	_ = Save(client, index, expected)

	// Fetch saved item
	resp, _ := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	t.Logf("%s", resp.Source)

	var actual engine.Item
	_ = json.Unmarshal(resp.Source, &actual)
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	// Verify result
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
