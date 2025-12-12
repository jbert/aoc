package y2025

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"github.com/jbert/aoc/graph"
	"github.com/jbert/aoc/matrix"
	"github.com/jbert/aoc/year"
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
	vorder := g.Vertices()
	slices.Sort(vorder)
	order := make(map[string]int)
	for i, v := range vorder {
		order[v] = i
	}
	m := g.AdjacencyMatrixFromOrder(order)
	// fmt.Printf("m %v\n", m)
	// for i := range 2 {
	// 	for j := range 2 {
	// 		fmt.Printf("i %d j %d Get(i,j) %d\n", i, j, m.Get(i, j))
	// 	}
	// }
	mx := len(order)
	ifr := order[fr]
	ito := order[to]

	// fmt.Printf("order %v\n", order)
	// fmt.Printf("fr %s to %s\n", fr, to)
	// fmt.Printf("ifr %d ito %d\n", ifr, ito)

	npaths := m.Get(ito, ifr)
	// fmt.Printf("1-length: %d\n", npaths)
	a := m.Copy()
	steps := 1
	for range mx {
		var err error
		a, err = matrix.Mult(a, m)
		if err != nil {
			panic("wtf")
		}
		steps++
		// fmt.Printf("a %v\n", a)
		npaths += a.Get(ito, ifr)
		// fmt.Printf("%d length: %v\n", steps, a.Get(ifr, ito))
	}
	return npaths
}
