package persist

import (
	"context"
	"github.com/olivere/elastic/v7"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			save(item)
		}
	}()
	return out
}

func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.199.137:9200"),
		elastic.SetSniff(false),
	)

	if err != nil {
		return "", err
	}

	resp, err := client.Index().
		Index("dating_profile").
		Type("zhenai").
		BodyJson(item).
		Do(context.Background())

	if err != nil {
		return "", err
	}

	return resp.Id, nil
}