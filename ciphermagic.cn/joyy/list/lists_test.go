package list

import (
	"fmt"
	"strings"
	"testing"
)

func TestList(t *testing.T) {
	raw := []*Person{
		{Name: "aaa", Men: []Man{{ManName: "1"}, {ManName: "2"}}},
		{Name: "bbb", Men: []Man{{ManName: "3"}, {ManName: "4"}}},
	}
	res, ok := Lists[Person](raw).Flat(func(p any) []any {
		return Lists[Woman](p.(Person).Men).Map(func(m any) any {
			return Woman{WomanName: p.(Person).Name + "-" + m.(Man).ManName}
		}).ToList()
	}).Filter(func(a any) bool {
		return strings.HasPrefix(a.(Woman).WomanName, "aaa")
	}).Map(func(a any) any {
		return Person{Name: a.(Woman).WomanName}
	}).Max(func(i, j any) bool {
		return i.(Person).Name > j.(Person).Name
	}).FindFirst()
	fmt.Println(res, ok)
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
