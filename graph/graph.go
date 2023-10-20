package graph

import (
	"fmt"
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
