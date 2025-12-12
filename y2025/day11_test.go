package y2025

import (
	"fmt"
	"testing"

	"github.com/jbert/aoc/graph"
	"github.com/stretchr/testify/assert"
)

func TestCountPaths(t *testing.T) {
	a := assert.New(t)

	tcs := []struct {
		g        *graph.Graph[string]
		expected int
	}{{graph.NewFromEdges([]graph.Edge[string]{
		{"A", "X", 0},
	}, false),
		1,
	},
		{graph.NewFromEdges([]graph.Edge[string]{
			{"A", "X", 0},
			{"A", "B", 0},
			{"B", "X", 0},
		}, false),
			2,
		},
		{graph.NewFromEdges([]graph.Edge[string]{
			{"A", "X", 0},
			{"A", "B", 0},
			{"B", "X", 0},
			{"B", "C", 0},
			{"C", "X", 0},
		}, false),
			3,
		}}

	for i, tc := range tcs {
		fmt.Printf("G %+v\n", tc.g)
		got := countPaths("A", "X", tc.g)
		a.Equal(tc.expected, got, fmt.Sprintf("%d: adj", i))
	}
}
