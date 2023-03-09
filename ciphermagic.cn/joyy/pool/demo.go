package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"math"
	"runtime"
	"time"
)

func main() {
	liveWithTask()
}

// -============================ dead =============================

func dead() {
	for i := 0; i < math.MaxInt64; i++ {
		go func(i int) {
			fmt.Println("goroutine count = ", runtime.NumGoroutine())
		}(i)
	}
}

// -============================ wantToDie =============================

func wantToDie() {
	for i := 0; i < math.MaxInt64; i++ {
		go func(i int) {
			time.Sleep(5 * time.Second)
		}(i)
	}
}

// -============================ live with channel =============================

var channel = make(chan bool, 3)

func liveWithChan() {
	for i := 0; i < math.MaxInt64; i++ {
		channel <- true
		go func(i int) {
			defer func() { <-channel }()
			fmt.Println("goroutine count = ", runtime.NumGoroutine())

		}(i)
	}
}

// -========================== live with semaphore =============================

func liveWithSemaphore() {
	sem := semaphore.NewWeighted(3)
	for i := 0; i < math.MaxInt64; i++ {
		sem.Acquire(context.Background(), 1)
		go func(i int) {
			defer func() { sem.Release(1) }()
			fmt.Println("goroutine count = ", runtime.NumGoroutine())
		}(i)
	}
}

// -========================== live with task =============================

var ch = make(chan int)

func sendTask(task int, ch chan int) {
	ch <- task
}

func work(ch chan int) {
	for range ch {
		fmt.Println("goroutine count = ", runtime.NumGoroutine())
	}
}

func liveWithTask() {
	// 任务消费者
	for i := 0; i < 3; i++ {
		go work(ch)
	}

	// 任务生产者
	for i := 0; i < math.MaxInt64; i++ {
		sendTask(i, ch)
	}
}
