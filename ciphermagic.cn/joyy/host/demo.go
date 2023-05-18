package main

import "fmt"

func main() {
	fmt.Println(HostInfo.Idc)
	fmt.Println(HostInfo.Name)
	var i *int
	fmt.Println(*i == 1)
}
