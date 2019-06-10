package main

import "fmt"

// 这是一个并行计算的例子，由两个goroutine进行并行的累加计算，待这两个计算过程都完成后打印计算结果
func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	resultChan := make(chan int)
	go sum(values[:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)
	sum1, sum2 := <-resultChan, <-resultChan // 接收结果

	fmt.Printf("Result: %d + %d = %d", sum1, sum2, sum1+sum2)
}

func sum(values []int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	resultChan <- sum // 将计算结果发送到channel中
}
