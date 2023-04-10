package list

import (
	"fmt"
	"reflect"
	"sort"
)

type List[T any] struct {
	list []any
}

func Lists[T any](items any) *List[T] {
	rv := reflect.ValueOf(items)
	if rv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("not supported type: %v, please use slice instead", rv.Kind()))
	}
	l := rv.Len()
	s := make([]any, 0, l)
	for i := 0; i < l; i++ {
		s = append(s, rv.Index(i).Interface())
	}
	return &List[T]{
		list: s,
	}
}

func (s *List[T]) Filter(fn func(any) bool) *List[T] {
	l := make([]any, 0)
	for _, e := range s.list {
		if fn(e) {
			l = append(l, e)
		}
	}
	s.list = l
	return s
}

func (s *List[T]) Map(fn func(any) any) *List[T] {
	l := make([]any, 0)
	for _, element := range s.list {
		l = append(l, fn(element))
	}
	return &List[T]{
		list: l,
	}
}

func (s *List[T]) Flat(fn func(any) []any) *List[T] {
	l := make([]any, 0)
	for _, element := range s.list {
		l = append(l, fn(element)...)
	}
	return &List[T]{
		list: l,
	}
}

func (s *List[T]) Sort(fn func(i, j any) bool) *List[T] {
	if len(s.list) <= 0 {
		return s
	}
	sort.SliceStable(s.list, func(i, j int) bool {
		return fn(s.list[i], s.list[j])
	})
	return s
}

func (s *List[T]) Max(fn func(i, j any) bool) (T, bool) {
	return s.Sort(fn).FindFirst()
}

func (s *List[T]) FindFirst() (T, bool) {
	if len(s.list) <= 0 {
		var nonsense T
		return nonsense, false
	}
	return s.list[0].(T), true
}

func (s *List[T]) ToList() []any {
	return s.list
}

func (s *List[T]) Collect() []T {
	t := make([]T, 0)
	for _, a := range s.list {
		t = append(t, a.(T))
	}
	return t
}
