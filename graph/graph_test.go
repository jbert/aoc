package graph

import (
	"sort"
	"testing"

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
}
