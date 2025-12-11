package y2025

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jbert/aoc/graph"
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
	// for _, p := range paths {
	// fmt.Printf("%v\n", p)
	// }
	fmt.Printf("Part 1: %d\n", len(paths))
	// fmt.Printf("Part 2: %d\n", fun.Sum(joltages))

	return nil
}
