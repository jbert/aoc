package graph

import (
	"fmt"
	"sort"
	"testing"

	"github.com/jbert/aoc/matrix"
	"github.com/jbert/fun"
	"github.com/stretchr/testify/assert"
)

func TestVertices(t *testing.T) {
	a := assert.New(t)

	edges := []Edge[string]{
		{"A", "B", 0},
		{"B", "C", 0},
		{"B", "D", 0},
		{"C", "E", 0},
		{"D", "E", 0},
	}
	g := NewFromEdges(edges, true)
	vs := g.Vertices()
	sort.Strings(vs)
	a.Equal([]string{"A", "B", "C", "D", "E"}, vs, "Correct vertices")

	es := g.Edges()
	ess := fun.Map(func(e Edge[string]) string { return fmt.Sprintf("%s-%s", e.From, e.To) }, es)
	sort.Strings(ess)
	expected := fun.Flatten(fun.Map(func(s string) []string {
		return []string{s, string(fun.Reverse([]byte(s)))}
	}, []string{"A-B", "B-C", "B-D", "C-E", "D-E"}))
	sort.Strings(expected)
	a.Equal(expected, ess)

	g2, _ := g.Remove("B")
	vs = g2.Vertices()
	sort.Strings(vs)
	a.Equal([]string{"A", "C", "D", "E"}, vs, "Correct vertices")

	es = g2.Edges()
	ess = fun.Map(func(e Edge[string]) string { return fmt.Sprintf("%s-%s", e.From, e.To) }, es)
	sort.Strings(ess)
	expected = fun.Flatten(fun.Map(func(s string) []string {
		return []string{s, string(fun.Reverse([]byte(s)))}
	}, []string{"C-E", "D-E"}))
	sort.Strings(expected)
	a.Equal(expected, ess)
}

func TestRemoveAddEdge(t *testing.T) {
	a := assert.New(t)

	edges := []Edge[string]{
		{"A", "B", 0},
		{"B", "C", 0},
		{"B", "D", 0},
		{"C", "E", 0},
		{"D", "E", 0},
	}
	g := NewFromEdges(edges, false)
	a.Equal(5, len(g.Edges()))
	err := g.RemoveEdge("B", "C")
	a.Equal(4, len(g.Edges()))
	a.Nil(err, "no error")
	err = g.RemoveEdge("B", "C")
	a.Equal(4, len(g.Edges()))
	a.Equal(ErrNotFound, err, "correct error")

	g.AddEdge(Edge[string]{"B", "C", 1.0})
	a.Equal(5, len(g.Edges()))
	g.AddEdge(Edge[string]{"B", "C", 1.0})
	a.Equal(5, len(g.Edges()))
}

func TestAdjMatrix(t *testing.T) {
	a := assert.New(t)

	edges := []Edge[string]{
		{"A", "B", 0},
		{"B", "C", 0},
		{"B", "D", 0},
		{"C", "E", 0},
		{"D", "E", 0},
	}

	tcs := []struct {
		g        *Graph[string]
		vorder   []string
		expected matrix.Mat
	}{{NewFromEdges([]Edge[string]{
		{"A", "B", 0},
	}, false),
		[]string{"A", "B"},
		matrix.NewFromRows([][]int{
			{0, 1},
			{0, 0},
		})},
		{NewFromEdges([]Edge[string]{
			{"A", "B", 0},
			{"A", "C", 0},
			{"B", "C", 0},
		}, false),
			[]string{"A", "B", "C"},
			matrix.NewFromRows([][]int{
				{0, 1, 1},
				{0, 0, 1},
				{0, 0, 0},
			})},
		{NewFromEdges(edges, false),
			[]string{"A", "B", "C", "D", "E"},
			matrix.NewFromRows([][]int{
				{0, 1, 0, 0, 0},
				{0, 0, 1, 1, 0},
				{0, 0, 0, 0, 1},
				{0, 0, 0, 0, 1},
				{0, 0, 0, 0, 0},
			}),
		},
	}

	for i, tc := range tcs {
		order := make(map[string]int)
		for i, v := range tc.vorder {
			order[v] = i
		}
		got := tc.g.AdjacencyMatrixFromOrder(order)
		a.Equal(tc.expected, got, fmt.Sprintf("%d: adj", i))

	}
}
