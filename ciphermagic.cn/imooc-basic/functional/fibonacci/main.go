package main

import (
	"bufio"
	"ciphermagic.cn/imooc-basic/functional/fibonacci/fib"
	"fmt"
	"io"
	"strings"
)

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 1000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	// TODO: incorrect if p is to small
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	var f intGen = fib.Fibonacci()
	printFileContents(f)
}
