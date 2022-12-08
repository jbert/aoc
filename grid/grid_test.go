package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrid(t *testing.T) {
	a := assert.New(t)

	g := New[int](3, 4)
	a.Equal(3, g.Width(), "correct width")
	a.Equal(4, g.Height(), "correct height")

	a.Equal(0, g.Get(0, 0), "can read 0,0")
	g.Set(0, 0, 3)
	a.Equal(3, g.Get(0, 0), "can read set val at 0,0")

	a.Equal(0, g.Get(1, 2), "can read 1,2")
	g.Set(1, 2, 5)
	a.Equal(5, g.Get(1, 2), "can read set val at 1,2")

	t.Logf("g is %v\n", g)
}

func TestForEach(t *testing.T) {
	a := assert.New(t)

	g := New[int](2, 3)
	g.ForEach(func(i, j int) {
		g.Set(i, j, i)
	})
	expected := NewFromRows([][]int{{0, 1}, {0, 1}, {0, 1}})
	a.Equal(expected, g, "Can init with i")

	g.ForEach(func(i, j int) {
		g.Set(i, j, j)
	})
	expected = NewFromRows([][]int{{0, 0}, {1, 1}, {2, 2}})
	a.Equal(expected, g, "Can init with j")

	g.ForEach(func(i, j int) {
		g.Set(i, j, i+j)
	})
	expected = NewFromRows([][]int{{0, 1}, {1, 2}, {2, 3}})
	a.Equal(expected, g, "Can init with i+j")
}

func TestCombine(t *testing.T) {
	a := assert.New(t)

	g := New[int](2, 3)
	g.ForEach(func(i, j int) {
		g.Set(i, j, i)
	})

	h := New[int](2, 3)
	h.ForEach(func(i, j int) {
		h.Set(i, j, j)
	})
	t.Logf("G is %v\n", g)
	t.Logf("H is %v\n", h)

	got := g.Combine(h, func(a, b int) int { return a + b })
	expected := NewFromRows([][]int{{0, 1}, {1, 2}, {2, 3}})
	a.Equal(expected, got, "Can add two grids with combine")
}
