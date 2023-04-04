package list

import (
	"fmt"
	"reflect"
)

type List[T any] struct {
	list []any
}

func Lists[T any](items any) *List[T] {
	rv := reflect.ValueOf(items)
	switch rv.Kind() {
	case reflect.Slice:
	default:
		panic(fmt.Sprintf("not supported type: %v, please use slice intead", rv.Kind()))
	}
	l := rv.Len()
	s := make([]any, 0, l)
	for i := 0; i < l; i++ {
		_v := rv.Index(i)
		switch _v.Kind() {
		case reflect.Pointer:
			s = append(s, rv.Index(i).Elem().Interface())
		default:
			s = append(s, rv.Index(i).Interface())
		}
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
