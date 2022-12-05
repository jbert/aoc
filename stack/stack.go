package stack

import (
	"fmt"
	"strings"
)

type Stack[A any] []A

func New[A any]() Stack[A] {
	return make([]A, 0)
}

func (s Stack[A]) String() string {
	b := &strings.Builder{}
	first := true
	for _, a := range s {
		if first {
			first = false
		} else {
			b.WriteString(",")
		}
		b.WriteString(fmt.Sprintf("%v", a))
	}
	return b.String()
}

func (s *Stack[A]) Push(a A) {
	*s = append(*s, a)
}

func (s *Stack[A]) MustPeek() A {
	return (*s)[len(*s)-1]
}

func (s *Stack[A]) MustPop() A {
	a := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return a
}

func (s *Stack[A]) Pop() (A, bool) {
	if len(*s) == 0 {
		var a A
		return a, false
	}
	return s.MustPop(), true
}

func (s *Stack[A]) Size() int {
	return len(*s)
}

func (s *Stack[A]) Peek() (A, bool) {
	if len(*s) == 0 {
		var a A
		return a, false
	}
	return s.MustPeek(), true
}
