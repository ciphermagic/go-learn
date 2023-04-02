package list

type List struct {
	list []any
}

func Lists[T any](items []T) *List {
	s := make([]any, 0)
	for _, item := range items {
		s = append(s, item)
	}
	return &List{
		list: s,
	}
}

func (s *List) Filter(fn func(any) bool) *List {
	l := make([]any, 0)
	for _, e := range s.list {
		if fn(e) {
			l = append(l, e)
		}
	}
	s.list = l
	return s
}

func (s *List) Map(fn func(any) any) *List {
	l := make([]any, 0)
	for _, element := range s.list {
		l = append(l, fn(element))
	}
	return &List{
		list: l,
	}
}

func (s *List) Flat(fn func(any) []any) *List {
	l := make([]any, 0)
	for _, element := range s.list {
		l = append(l, fn(element)...)
	}
	return &List{
		list: l,
	}
}

func (s *List) ToList() []any {
	return s.list
}

func (s *List) Collect(f func(any) any) []any {
	target := make([]any, 0)
	for _, v := range s.list {
		target = append(target, f(v))
	}
	return target
}
