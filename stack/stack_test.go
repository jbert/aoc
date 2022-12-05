package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	a := assert.New(t)

	s := New[int]()
	_, ok := s.Pop()
	a.False(ok, "Can't pop empty stack")

	s.Push(2)
	s.Push(4)
	s.Push(6)
	a.Equal(3, s.Size(), "Pushed 3 elts")

	x, ok := s.Peek()
	a.True(ok, "can peek stack")
	a.Equal(6, x, "Top of stack is correct")
	a.Equal(3, s.Size(), "Peek doesn't change size")

	x, ok = s.Pop()
	a.True(ok, "can pop stack")
	a.Equal(6, x, "Top of stack is correct")
	a.Equal(2, s.Size(), "Pop does change size")

	x, ok = s.Pop()
	a.True(ok, "can pop stack again")
	a.Equal(4, x, "Top of stack is correct")
	a.Equal(1, s.Size(), "Pop does change size")
}
