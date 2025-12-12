package y2025

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jbert/aoc/astar"
	"github.com/jbert/aoc/graph"
	"github.com/jbert/aoc/year"
	"github.com/jbert/fun"
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
	// for _, p := range paths {
	// fmt.Printf("%v\n", p)
	// }
	fmt.Printf("Part 1: %d\n", len(paths))

	/*
		paths = g.FindAllPaths("svr", "out")
		paths = fun.Filter(func(p graph.Path[string]) bool { return p.Contains("dac") }, paths)
		paths = fun.Filter(func(p graph.Path[string]) bool { return p.Contains("fft") }, paths)
		fmt.Printf("Part 2: %d\n", len(paths))
	*/
	/*
			fmt.Printf("JB1\n")
			s2d := g.FindAllPaths("svr", "dac")
			fmt.Printf("JB2\n")
			d2f := g.FindAllPaths("dac", "fft")
			fmt.Printf("JB3\n")
			f2o := g.FindAllPaths("fft", "out")
			fmt.Printf("JB4\n")

			s2f := g.FindAllPaths("svr", "fft")
			fmt.Printf("JB5\n")
			f2d := g.FindAllPaths("fft", "dac")
			fmt.Printf("JB6\n")
			d2o := g.FindAllPaths("dac", "out")
			fmt.Printf("JB7\n")

		count := len(s2d)*len(d2f)*len(f2o) + len(s2f)*len(f2d)*len(d2o)
		fmt.Printf("Part 2: %d\n", count)
	*/

	// There is a path from fft->dac, but not one from dac->fft

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
	paths := getPaths(fr, to, *g, 0)
	return len(paths)
}

func pathToStr(p graph.Path[string]) string {
	return strings.Join(p, "-")
}

func strToPath(s string) graph.Path[string] {
	return strings.Split(s, "-")
}

func getPaths(fr string, to string, g graph.Graph[string], depth int) []graph.Path[string] {
	fmt.Printf("GP >>> %d\n", depth)
	pathStrs := set.New[string]()

	starPath, err := astar.Astar(fr, to, g, func(string) float64 { return 1.0 })
	if err != nil {
		return fun.Map(strToPath, pathStrs.ToList())
	}
	// if depth == 0 {
	fmt.Printf("%v\n", starPath)
	// }
	pathStrs.Insert(pathToStr(starPath))
	for i, eto := range starPath {
		if i == 0 {
			continue
		}
		efr := starPath[i-1]
		g.RemoveEdge(efr, eto)
		fmt.Printf("RM: %s -> %s\n", efr, eto)
		cPaths := getPaths(fr, to, g, depth+1)
		g.AddEdge(graph.Edge[string]{From: efr, To: eto})
		fmt.Printf("AD: %s -> %s\n", efr, eto)
		pathStrs = pathStrs.Union(set.NewFromList(fun.Map(pathToStr, cPaths)))
		if depth == 0 {
			fmt.Printf("%d\n", pathStrs.Size())
		}
	}
	fmt.Printf("GP <<< %d\n", depth)
	return fun.Map(strToPath, pathStrs.ToList())
}
