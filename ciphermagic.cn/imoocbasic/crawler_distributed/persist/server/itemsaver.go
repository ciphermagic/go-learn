package main

import (
	"ciphermagic.cn/imoocbasic/crawler/config"
	"ciphermagic.cn/imoocbasic/crawler_distributed/persist"
	"ciphermagic.cn/imoocbasic/crawler_distributed/rpcsupport"
	"flag"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.199.137:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
