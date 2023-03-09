package main

import (
	"fmt"
	"github.com/bytedance/gopkg/util/gopool"
	"time"
)

func main() {
	gopool.Go(func() {
		fmt.Println("aaa")
	})
	time.Sleep(time.Hour)
}
