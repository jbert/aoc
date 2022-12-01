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
