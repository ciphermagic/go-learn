package persist

import (
	"ciphermagic.cn/imoocbasic/crawler/model"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"testing"
)

func TestSaver(t *testing.T) {
	expected := model.Profile{
		Name:       "安静的雪",
		Gender:     "女",
		Age:        34,
		Height:     162,
		Weight:     57,
		Income:     "3001-5000元",
		Marriage:   "离异",
		Education:  "大学本科",
		Occupation: "人事/行政",
		Hokou:      "山东菏泽",
		Xinzuo:     "牡羊座",
		House:      "已购房",
		Car:        "未购车",
	}

	id, err := save(expected)

	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.199.137:9200"),
		elastic.SetSniff(false),
	)

	resp, err := client.Get().
		Index("dating_profile").
		Type("zhenai").
		Id(id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual model.Profile
	err = json.Unmarshal(resp.Source, &actual)

	if err != nil {
		panic(err)
	}

	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
