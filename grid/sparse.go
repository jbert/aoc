package grid

import (
	"github.com/jbert/aoc/pts"
)

// Sparse is a zero-indexed rectangular grid with Get/Set
// The zero value for A is return
type Sparse[A any] struct {
	minSet bool
	MinX   int
	MinY   int
	MaxX   int
	MaxY   int
	empty  A
	m      map[pts.P2]A
}

func NewSparse[A any](empty A) *Sparse[A] {
	return &Sparse[A]{
		empty: empty,
		m:     make(map[pts.P2]A),
	}
}

func (g Sparse[A]) GetPt(p pts.P2) A {
	v, ok := g.m[p]
	if ok {
		return v
	} else {
		return g.empty
	}
}

func (g *Sparse[A]) SetPt(p pts.P2, v A) {
	//	fmt.Printf("P: %s\n", p)
	g.m[p] = v
	g.updateMinMax(p)
}

func (g *Sparse[A]) updateMinMax(p pts.P2) {
	if !g.minSet || p.X < g.MinX {
		g.MinX = p.X
	}
	if p.X > g.MaxX {
		g.MaxX = p.X
	}
	if !g.minSet || p.Y < g.MinY {
		g.MinY = p.Y
	}
	if p.Y > g.MaxY {
		g.MaxY = p.Y
	}
	g.minSet = true
}
