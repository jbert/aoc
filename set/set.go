package set

import (
	"fmt"
	"strings"
)

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	return make(map[T]struct{})
}

func (s Set[T]) ForEach(f func(T)) {
	for k, _ := range s {
		f(k)
	}
}

func (s Set[T]) String() string {
	b := &strings.Builder{}
	first := true
	s.ForEach(func(e T) {
		if first {
			first = false
		} else {
			fmt.Fprintf(b, ", ")
		}
		fmt.Fprintf(b, "%v", e)
	})
	return b.String()
}

func Map[T, U comparable](s Set[T], f func(T) U) Set[U] {
	m := New[U]()
	s.ForEach(func(e T) {
		m.Insert(f(e))
	})
	return m
}

func (s Set[T]) Contains(a T) bool {
	_, ok := s[a]
	return ok
}

func (s Set[T]) Intersect(a Set[T]) Set[T] {
	i := New[T]()
	for x, _ := range s {
		if a.Contains(x) {
			i.Insert(x)
		}
	}
	return i
}

func (s Set[T]) Union(a Set[T]) Set[T] {
	u := New[T]()
	for x, _ := range s {
		u.Insert(x)
	}
	for x, _ := range a {
		u.Insert(x)
	}
	return u
}

func NewFromList[T comparable](as []T) Set[T] {
	s := New[T]()
	s.InsertList(as)
	return s
}

func (s Set[T]) Equals(a Set[T]) bool {
	if s.Size() != a.Size() {
		return false
	}
	for x, _ := range s {
		if !a.Contains(x) {
			return false
		}
	}
	return true
}

func (s Set[T]) InsertList(as []T) {
	for _, a := range as {
		s.Insert(a)
	}
}

func (s Set[T]) Insert(a T) {
	s[a] = struct{}{}
}

func (s Set[T]) Remove(a T) {
	delete(s, a)
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s Set[T]) ToList() []T {
	l := make([]T, 0)
	for k, _ := range s {
		l = append(l, k)
	}
	return l
}
