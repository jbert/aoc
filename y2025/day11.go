package y2025

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jbert/aoc/graph"
	"github.com/jbert/aoc/year"
	"github.com/jbert/set"
)

type Day11 struct{ year.Year }

func parseEdges(l string) []graph.Edge[string] {
	bits := strings.Split(l, ": ")
	node := bits[0]
	dests := strings.Split(bits[1], " ")
	var edges []graph.Edge[string]
	for _, d := range dests {
		edges = append(edges, graph.Edge[string]{From: node, To: d})
	}
	return edges
}

func (d *Day11) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")
	var edges []graph.Edge[string]
	for _, l := range lines {
		lineEdges := parseEdges(l)
		for _, e := range lineEdges {
			edges = append(edges, e)
		}
	}

	g := graph.NewFromEdges(edges, false)
	g.ToDot(os.Stdout, "reactor")

	paths := g.FindAllPaths("you", "out")
	fmt.Printf("Part 1: %d\n", len(paths))

	fmt.Printf("JB1\n")
	s2f := countPaths("svr", "fft", g.Copy())
	fmt.Printf("JB2: %d\n", s2f)
	f2d := countPaths("fft", "dac", g.Copy())
	fmt.Printf("JB3: %d\n", f2d)
	d2o := countPaths("dac", "out", g.Copy())
	fmt.Printf("s2f %d f2d %d d2o %d\n", s2f, f2d, d2o)
	fmt.Printf("Part 2: %d\n", s2f*f2d*d2o)

	return nil
}

func countPaths(fr string, to string, g *graph.Graph[string]) int {
	labels := labelVertices(fr, to, g)
	fmt.Printf("labels: %+v\n", labels)
	return labels[to]
}

func labelVertices(fr string, to string, g *graph.Graph[string]) map[string]int {
	labels := make(map[string]int)

	addTo := func(v string, n int) {
		nn := labels[v]
		nn++
		labels[v] = nn
	}
	addOne := func(v string) { addTo(v, 1) }
	todo := set.NewFromList[string](g.Neighbours(fr))
	todo.ForEach(addOne)

	for todo.Size() > 0 {
		v, err := todo.Take()
		if err != nil {
			panic("take on non-empty set failed")
		}
		ns := g.Neighbours(v)
		for _, nv := range ns {
			addTo(nv, labels[v])
			todo.Insert(nv)
		}
	}

	return labels
}
