package main

import (
	"fmt"
	"math/rand"
)

type RandomOrder struct {
	count    uint32
	coprimes []uint32
}

type RandomEnum struct {
	i     uint32
	count uint32
	pos   uint32
	inc   uint32
}

func NewRandomOrder(count int) RandomOrder {
	var ord RandomOrder
	ord.reset(uint32(count))
	return ord
}

func (ord *RandomOrder) reset(count uint32) {
	ord.count = count
	ord.coprimes = ord.coprimes[:0]
	for i := uint32(1); i <= count; i++ {
		if gcd(i, count) == 1 {
			ord.coprimes = append(ord.coprimes, i)
		}
	}
}

func (ord *RandomOrder) start() RandomEnum {
	i := rand.Uint32()
	return RandomEnum{
		count: ord.count,
		pos:   i % ord.count,
		inc:   ord.coprimes[i/ord.count%uint32(len(ord.coprimes))],
	}
}

func (enum *RandomEnum) done() bool {
	return enum.i == enum.count
}

func (enum *RandomEnum) next() {
	enum.i++
	enum.pos = (enum.pos + enum.inc) % enum.count
}

func (enum *RandomEnum) position() uint32 {
	return enum.pos
}

func gcd(a, b uint32) uint32 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	ord := NewRandomOrder(8)
	for i := 0; i <= 50; i++ {
		for enum := ord.start(); !enum.done(); enum.next() {
			fmt.Print(enum.position())
		}
		fmt.Println()
	}
}
