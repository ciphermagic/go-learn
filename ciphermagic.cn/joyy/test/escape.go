package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println("aaaaaaaaaaa")
		}
	}()
	select {}
}
