package main

import "fmt"

func main() {
	m := map[string]string{
		"name":   "learn",
		"course": "golang",
		"site":   "imooc",
	}

	m2 := make(map[string]int) // empty map

	var m3 map[string]int // nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)

	if nothing, ok := m["xxxx"]; ok {
		fmt.Println(nothing)
	} else {
		fmt.Println("Key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
