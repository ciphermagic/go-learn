package list

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	raw := []Person{
		{Name: "aaa", Men: []Man{{ManName: "1"}, {ManName: "2"}}},
		{Name: "bbb", Men: []Man{{ManName: "3"}, {ManName: "4"}}},
	}
	res := Lists(raw).
		Flat(func(i any) []any {
			return Lists(i.(Person).Men).ToList()
		}).
		Map(func(i any) any {
			return Woman{WomanName: i.(Man).ManName}
		}).
		Collect(func(i any) any {
			return i.(Woman)
		})
	fmt.Println(res)
}

type Person struct {
	Name string
	Men  []Man
}

type Man struct {
	ManName string
}

type Woman struct {
	WomanName string
}
