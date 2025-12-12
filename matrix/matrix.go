package matrix

import (
	"fmt"

	"github.com/jbert/aoc/grid"
)

type Mat = grid.Grid[int]

func New(w int, h int) Mat {
	return grid.New[int](w, h)
}

func NewFromRows(rows [][]int) grid.Grid[int] {
	return grid.NewFromRows(rows)
}

func Mult(a Mat, b Mat) (Mat, error) {
	if a.Width() != b.Height() || a.Height() != b.Width() {
		return nil, fmt.Errorf("incompatible: %dx%d and %dx%d", a.Width(), a.Height(), b.Width(), b.Height())
	}
	w := b.Width()
	h := a.Height()
	m := New(w, h)
	for i := range w {
		for j := range h {
			for k := range a.Width() {
				v := m.Get(i, j)
				m.Set(i, j, v+a.Get(k, j)*b.Get(i, k))
			}
		}
	}
	return m, nil
}
