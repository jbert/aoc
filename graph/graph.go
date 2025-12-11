package graph

import (
	"fmt"
	"io"
	"strings"

	"github.com/jbert/fun"
	"github.com/jbert/set"
)

type Edge[V comparable] struct {
	From   V
	To     V
	Weight float64
}

func (e Edge[V]) String() string {
	return fmt.Sprintf("%v - %v", e.From, e.To)
}

func (e Edge[V]) Reverse() Edge[V] {
	return Edge[V]{From: e.To, To: e.From}
}

type Graph[V comparable] map[V]set.Set[Edge[V]]

func (g Graph[V]) ToDot(w io.Writer, name string) {
	// TODO: support directed/undirected properly (in code and Dot)
	fmt.Fprintf(w, "graph %s {\n", name)
	edges := g.Edges()
	for _, e := range edges {
		fmt.Fprintf(w, "\t%v -- %v\n", e.From, e.To)
	}
	fmt.Fprintf(w, "}\n")
	return
}

func (g Graph[V]) Edges() []Edge[V] {
	var edges []Edge[V]
	for _, s := range g {
		s.ForEach(func(e Edge[V]) {
			edges = append(edges, e)
		})
	}
	return edges
}

func (g Graph[V]) Vertices() []V {
	var vs []V
	for v := range g {
		vs = append(vs, v)
	}
	return vs
}

func (g Graph[V]) String() string {
	b := &strings.Builder{}
	for v, s := range g {
		fmt.Fprintf(b, "%v:\t%v\n", v, s)
	}
	return b.String()
}

func NewFromEdges[V comparable](edges []Edge[V], undirected bool) *Graph[V] {
	g := Graph[V]{}
	for _, edge := range edges {
		g.addEdge(edge)
		if undirected {
			g.addEdge(edge.Reverse())
		}
	}
	return &g
}

func (g Graph[V]) addVertex(v V) {
	s, ok := g[v]
	if !ok {
		s = set.New[Edge[V]]()
	}
	g[v] = s
}

func (g Graph[V]) addEdge(e Edge[V]) {
	s, ok := g[e.From]
	if !ok {
		s = set.New[Edge[V]]()
	}
	s.Insert(e)
	g[e.From] = s
}

func (g Graph[V]) Neighbours(v V) []V {
	edges := g[v].ToList()
	return fun.Map(func(e Edge[V]) V { return e.To }, edges)
}

func (g Graph[V]) IsVertex(v V) bool {
	_, ok := g[v]
	return ok
}

func (g Graph[V]) Weight(from, to V) float64 {
	s, ok := g[from]
	if !ok {
		panic(fmt.Sprintf("Request for weight of non-existent edge from [%v]", from))
	}
	var weight float64
	weightFound := false
	s.ForEach(func(e Edge[V]) {
		if e.To == to {
			weight = e.Weight
			weightFound = true
		}
	})
	if !weightFound {
		panic(fmt.Sprintf("Request for weight of non-existent edge to [%v]", to))
	}
	return weight
}

func (g Graph[V]) Remove(v V) (*Graph[V], []Edge[V]) {
	edges := g.Edges()
	removedEdges := fun.Filter(func(e Edge[V]) bool {
		return e.From == v || e.To == v
	}, edges)
	keepEdges := fun.Filter(func(e Edge[V]) bool {
		return !(e.From == v || e.To == v)
	}, edges)

	g2 := NewFromEdges(keepEdges, false)
	ws := fun.Filter(func(w V) bool { return v != w }, g.Vertices())
	for _, w := range ws {
		g2.addVertex(w)
	}

	return g2, removedEdges
}

type Path[V comparable] []Edge[V]

func (p Path[V]) Prepend(e Edge[V]) Path[V] {
	return append(Path[V]{e}, p...)
}

func (g Graph[V]) FindAllPaths(fr V, to V) []Path[V] {
	if fr == to {
		return []Path[V]{{}}
	}

	steps := g[fr]
	var allPaths []Path[V]
	steps.ForEach(func(e Edge[V]) {
		rest := g.FindAllPaths(e.To, to)
		paths := fun.Map(func(p Path[V]) Path[V] { return p.Prepend(e) }, rest)
		for _, p := range paths {
			allPaths = append(allPaths, p)
		}
	})
	return allPaths
}
