package grid

import "github.com/jbert/aoc/pts"

// Grid is a zero-indexed rectangular grid with Get/Set
type Grid[A any] [][]A

func NewFromRows[A any](rows [][]A) Grid[A] {
	return rows
}

func New[A any](w, h int) Grid[A] {
	g := make([][]A, h)
	for y := range g {
		g[y] = make([]A, w)
	}
	return g
}

func (g Grid[A]) Rows() [][]A {
	return g
}

func (g Grid[A]) Row(y int) []A {
	return g[y]
}

func (g Grid[A]) GetPt(p pts.P2) A {
	return g.Get(p.X, p.Y)
}

func (g Grid[A]) SetPt(p pts.P2, v A) {
	g.Set(p.X, p.Y, v)
}

// CardinalNeighbourPts are up/down/left/right pts inside the grid
func (g Grid[A]) CardinalNeighbourPts(p pts.P2) []pts.P2 {
	var np []pts.P2
	if p.X > 0 {
		np = append(np, pts.P2{p.X - 1, p.Y})
	}
	if p.Y > 0 {
		np = append(np, pts.P2{p.X, p.Y - 1})
	}
	if p.X < g.Width()-1 {
		np = append(np, pts.P2{p.X + 1, p.Y})
	}
	if p.Y < g.Height()-1 {
		np = append(np, pts.P2{p.X, p.Y + 1})
	}
	return np
}

func (g Grid[A]) Get(x, y int) A {
	return g[y][x]
}

func (g Grid[A]) Set(x, y int, v A) {
	g[y][x] = v
}

func (g Grid[A]) Width() int {
	if len(g) == 0 {
		return 0
	}
	return len(g[0])
}

func (g Grid[A]) Height() int {
	return len(g)
}

func (g Grid[A]) ForEachVal(f func(A)) {
	g.ForEach(func(i, j int) {
		f(g.Get(i, j))
	})
}

func (g Grid[A]) ForEach(f func(int, int)) {
	for j, row := range g {
		for i := range row {
			f(i, j)
		}
	}
}

func (g Grid[A]) ForEachV(f func(int, int, A)) {
	for j, row := range g {
		for i := range row {
			v := g.Get(i, j)
			f(i, j, v)
		}
	}
}

func (g Grid[A]) Combine(h Grid[A], f func(A, A) A) Grid[A] {
	r := New[A](g.Width(), g.Height())
	r.ForEach(func(i, j int) {
		gVal := g.Get(i, j)
		hVal := h.Get(i, j)
		r.Set(i, j, f(gVal, hVal))
	})
	return r
}
