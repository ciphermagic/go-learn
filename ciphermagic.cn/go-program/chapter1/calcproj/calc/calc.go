package main

import (
	"fmt"
	"os"
	"strconv"

	"ciphermagic.cn/go-program/chapter1/calcproj/simplemath"
)

var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe commands are:\nadd\tAddition of two values.")
}

func main() {
	args := os.Args
	if args == nil || len(args) < 3 {
		Usage()
		return
	}

	switch args[1] {
	case "add":
		if len(args) != 4 {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("Result: ", ret)
	default:
		Usage()
	}
}
