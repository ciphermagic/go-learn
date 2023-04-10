package list

import (
	"testing"
)

var orders = []Order{
	{
		Id: "o1",
		Products: []Product{
			{
				Id:    "p1",
				Price: 1,
			},
			{
				Id:    "p2",
				Price: 2,
			},
		},
	},
	{
		Id: "o2",
		Products: []Product{
			{
				Id:    "p3",
				Price: 3,
			},
			{
				Id:    "p4",
				Price: 4,
			},
		},
	},
}

func TestFilter(t *testing.T) {
	res := Lists[Order](orders).Filter(func(o any) bool {
		return o.(Order).Id == "o2"
	}).Collect()
	t.Log(res) // [{o2 [{p3 3} {p4 4}]}]
}

func TestMap(t *testing.T) {
	res := Lists[CustomOrder](orders).Map(func(o any) any {
		return CustomOrder{
			Id: "custom-" + o.(Order).Id,
		}
	}).Collect()
	t.Log(res) // [{custom-o1} {custom-o2}]
}

func TestFlatAndMap(t *testing.T) {
	res := Lists[CustomOrder](orders).
		Flat(func(o any) []any {
			return Lists[any](o.(Order).Products).ToList()
		}).
		Map(func(p any) any {
			return CustomOrder{
				Id: "ProductId-" + p.(Product).Id,
			}
		}).Collect()
	t.Log(res) // [{ProductId-p1} {ProductId-p2} {ProductId-p3} {ProductId-p4}]
}

func TestMax(t *testing.T) {
	res, found := Lists[Product](orders).
		Flat(func(o any) []any {
			return Lists[any](o.(Order).Products).ToList()
		}).
		Max(func(i, j any) bool {
			return i.(Product).Price > j.(Product).Price
		})
	t.Log(found, res) // true {p4 4}
}

type Order struct {
	Id       string
	Products []Product
}

type Product struct {
	Id    string
	Price int
}

type CustomOrder struct {
	Id string
}
