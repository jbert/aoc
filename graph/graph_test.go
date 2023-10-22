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
	a.Equal(vs, []string{"A", "B", "C", "D", "E"}, "Correct vertices")

	es := g.Edges()
	ess := fun.Map(func(e Edge[string]) string { return fmt.Sprintf("%s-%s", e.From, e.To) }, es)
	sort.Strings(ess)
	fmt.Printf("ess %v\n", ess)
	expected := fun.Flatten(fun.Map(func(s string) []string {
		return []string{s, string(fun.Reverse([]byte(s)))}
	}, []string{"A-B", "B-C", "B-D", "C-E", "D-E"}))
	sort.Strings(expected)
	a.Equal(ess, expected)
}
