package grid

import (
	"github.com/jbert/aoc/pts"
	"github.com/jbert/fun"
)

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

func NewFromFunc[A any](w, h int, f func(pts.P2) A) Grid[A] {
	g := make([][]A, h)
	for y := range g {
		g[y] = make([]A, w)
		for x := range g[y] {
			g[y][x] = f(pts.P2{x, y})
		}
	}
	return g
}

/*
func (g Grid[A]) Reorder(order []int) Grid[A] {
	ng := New[A](g.Width(), g.Height())
	if len(order) != g.Width() {
		panic("wtf")
	}
	for i := range g.Width() {
		ii := order[i]
		for j := range g.Height() {
			jj := order[j]
			ng[jj][ii] = g[j][i]
		}
	}
	return ng
}
*/

func (g Grid[A]) Rows() [][]A {
	return g
}

func (g Grid[A]) Copy() Grid[A] {
	return NewFromRows(g.Rows())
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

func (g Grid[A]) Contains(p pts.P2) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < g.Width() && p.Y < g.Height()
}

// CardinalNeighbourPts are up/down/left/right pts inside the grid
func (g Grid[A]) CardinalNeighbourPts(p pts.P2) []pts.P2 {
	return fun.Filter(g.Contains, fun.Map(p.Add, pts.NESW))
}

// The (up to) 8 neighbouring points
func (g Grid[A]) AllNeighbourPts(p pts.P2) []pts.P2 {
	return fun.Filter(g.Contains, fun.Map(p.Add, pts.NEIGHBOURS))
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
	g.ForEach(func(p pts.P2) {
		f(g.GetPt(p))
	})
}

func (g Grid[A]) ForEach(f func(p pts.P2)) {
	for j, row := range g {
		for i := range row {
			p := pts.P2{i, j}
			f(p)
		}
	}
}

func (g Grid[A]) ForEachV(f func(pts.P2, A)) {
	for j, row := range g {
		for i := range row {
			p := pts.P2{i, j}
			v := g.GetPt(p)
			f(p, v)
		}
	}
}

func (g Grid[A]) Combine(h Grid[A], f func(A, A) A) Grid[A] {
	r := New[A](g.Width(), g.Height())
	r.ForEach(func(p pts.P2) {
		gVal := g.GetPt(p)
		hVal := h.GetPt(p)
		r.SetPt(p, f(gVal, hVal))
	})
	return r
}

func Fmap[A, B any](ga Grid[A], f func(A) B) Grid[B] {
	return NewFromFunc(ga.Width(), ga.Height(), func(p pts.P2) B {
		a := ga.GetPt(p)
		return f(a)
	})
}
