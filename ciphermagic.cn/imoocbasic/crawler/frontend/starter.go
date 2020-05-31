package main

import (
	"ciphermagic.cn/imoocbasic/crawler/frontend/controller"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("imoocbasic/crawler/frontend/view")))
	http.Handle("/search", controller.CreateSearchResultHandler("imoocbasic/crawler/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
