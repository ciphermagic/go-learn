package main

import (
	"ciphermagic.cn/imooc-basic/crawler/frontend/controller"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("imooc-basic/crawler/frontend/view")))
	http.Handle("/search", controller.CreateSearchResultHandler("imooc-basic/crawler/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
