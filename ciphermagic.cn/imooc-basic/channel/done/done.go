package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, w worker) {
	for n := range w.data {
		fmt.Printf("worker %d received %c\n", id, n)
		w.done()
	}
}

type worker struct {
	data chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		make(chan int),
		func() {
			wg.Done()
		},
	}
	go doWorker(id, w)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup
	wg.Add(20)

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	for i, worker := range workers {
		worker.data <- 'a' + i
	}
	for i, worker := range workers {
		worker.data <- 'A' + i
	}

	wg.Wait()
}

func main() {
	chanDemo()
}
