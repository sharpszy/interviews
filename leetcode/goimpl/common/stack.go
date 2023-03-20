package common

type Stack []byte

func NewStack(cap ...uint) *Stack {
	size := 0
	if len(cap) > 0 {
		size = int(cap[0])
	}
	var s Stack = make([]byte, 0, size)
	return &s
}

func (s *Stack) Pop() (byte, bool) {
	l := s.Len()
	if l == 0 {
		return ' ', false
	}
	v := (*s)[l-1]
	*s = (*s)[:l-1]
	return v, true
}

func (s *Stack) Push(n ...byte) {
	*s = append(*s, n...)
}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Cap() int {
	return cap(*s)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}
