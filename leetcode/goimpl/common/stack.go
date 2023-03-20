package common

type Stack[T any] []T

func NewStack[T any](cap ...uint) *(Stack[T]) {
	size := 0
	if len(cap) > 0 {
		size = int(cap[0])
	}
	var s Stack[T] = make([]T, 0, size)
	return &s
}

func (s *Stack[T]) Pop() (v T, f bool) {
	if s.IsEmpty() {
		return v, false
	}
	v = (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return v, true
}

func (s *Stack[T]) Push(n ...T) {
	*s = append(*s, n...)
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) Cap() int {
	return cap(*s)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Len() == 0
}
