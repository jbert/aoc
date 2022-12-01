package astar

import (
	"testing"

	"github.com/jbert/aoc/graph"
	"github.com/stretchr/testify/assert"
)

func TestAstar(t *testing.T) {
	a := assert.New(t)
	// We should find the path which goes through 3, not through 2
	edges := []graph.Edge[int]{
		{1, 2, 2.0},
		{1, 3, 1.0},
		{2, 4, 1.0},
		{3, 4, 1.0},
	}
	g := graph.NewFromEdges(edges, false)

	one := func(v int) float64 { return 1.0 }
	gg := (Graph[int])(g)
	path, err := Astar(1, 4, gg, one)
	a.NoError(err, "No error in search")
	a.Equal([]int{1, 3, 4}, path)
}
