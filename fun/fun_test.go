package fun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	a := assert.New(t)

	a.Equal(Reverse([]int{1, 2, 3}), []int{3, 2, 1})
	a.Equal(Reverse([]int{1, 2, 3, 4}), []int{4, 3, 2, 1})
	a.Equal(Reverse([]int{1}), []int{1})
	a.Equal(Reverse([]int{}), []int{})

	a.Equal(Reverse([]string{"foo", "bar", "baz"}), []string{"baz", "bar", "foo"})
	a.Equal(Reverse([]string{}), []string{})
}

func TestSplitBy(t *testing.T) {
	a := assert.New(t)

	l := Iota(1, 10)
	a.Equal(SplitBy(l, 3), [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, "split dropping 1 on end")
	a.Equal(SplitBy(l, 5), [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}, "exact split")
	a.Equal(SplitBy([]int{}, 5), [][]int{}, "empty list")
}

func TestFlatten(t *testing.T) {
	a := assert.New(t)

	l := [][]int{{2, 3, 4}, {5, 6}, {7}, {}, {1, 2, 3}}
	a.Equal([]int{2, 3, 4, 5, 6, 7, 1, 2, 3}, Flatten(l), "Can flatten list")
	a.Equal([]int{}, Flatten([][]int{}), "can flatten zero list")
	a.Equal([]int{}, Flatten([][]int{{}}), "can flatten zero zero list")
}
