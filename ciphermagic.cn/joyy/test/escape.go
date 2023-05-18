package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	Consume()
	time.Sleep(3 * time.Second)
	close(q)
	time.Sleep(time.Hour)
}

var q = make(chan struct{})
var c = make(chan struct{})
var messageCh = make(chan string)

func Consume() {
	ctx := context.Background()
	go func() {
		for {
			fmt.Println("Consume")
			select {
			case <-q:
				fmt.Println("quit")
				close(c)
				return
			default:
				fmt.Println("receiving...")
				if msg, err := Receive(ctx); err == nil {
					fmt.Printf("received: %s\n", msg)
				}
			}
		}
	}()
}

func Receive(ctx context.Context) (string, error) {
	for {
		fmt.Println("Receive")
		select {
		case <-c:
			return "quit Receive", nil
		case cm, ok := <-messageCh:
			if !ok {
				return "", nil
			}
			return cm, nil
		case <-ctx.Done():
			return ctx.Err().Error(), ctx.Err()
		}
	}
}
