package main

import (
	"fmt"
	"time"

	"ciphermagic.cn/cipher/imoocbasic/interface/retriever/mock"
	"ciphermagic.cn/cipher/imoocbasic/interface/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func inspect(r Retriever) {
	fmt.Printf("Type => %T %v\n", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("inspecting => Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("inspecting => UserAgent", v.UserAgent)
	}
}

func main() {
	var r Retriever

	r = mock.Retriever{Contents: "this is a fake imooc.com"}
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Chrome/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	realRetriever := r.(*real.Retriever)
	fmt.Println("TimeOut:", realRetriever.TimeOut)

	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println("Contents:", mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	//fmt.Println(download(r))
}
