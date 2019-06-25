package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	personDB := make(map[string]PersonInfo)
	personDB["12345"] = PersonInfo{"12345", "Cipher", "GZ"}
	personDB["1"] = PersonInfo{"1", "Sunny", "HK"}
	person, ok := personDB["12345"]
	if ok {
		fmt.Println(person)
	} else {
		fmt.Println("Did not find person")
	}
}
