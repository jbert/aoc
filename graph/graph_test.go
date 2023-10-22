package graph

import (
	"fmt"
	"sort"
	"testing"

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
