package y2021

import (
	"fmt"
	"io"
	"strings"

	"github.com/jbert/aoc/fun"
	"github.com/jbert/aoc/graph"
)

type Day12 struct{ Year }

type Vertex string

func (v Vertex) IsSmall() bool {
	if v == "" {
		panic("Empty string vertex")
	}
	return string(v) == strings.ToLower(string(v))
}

type Visited struct {
	count      map[Vertex]int
	part       int
	smallTwice bool
}

func NewVisited(part int) *Visited {
	return &Visited{
		count: make(map[Vertex]int),
		part:  part,
	}
}

func (vst *Visited) Copy() *Visited {
	countCopy := make(map[Vertex]int)
	for v, c := range vst.count {
		countCopy[v] = c
	}
	vstCopy := Visited{
		count:      countCopy,
		part:       vst.part,
		smallTwice: vst.smallTwice,
	}
	return &vstCopy
}

func (vst *Visited) Visit(v Vertex) {
	vst.count[v]++
	if v.IsSmall() && vst.count[v] == 2 {
		vst.smallTwice = true
	}
}

func (vst Visited) CanVisit(v Vertex) bool {
	if v.IsSmall() {
		if v != "start" && v != "end" && vst.part == 2 && !vst.smallTwice {
			return true
		}
		return vst.count[v] == 0
	}
	return true
}

func CountPathsPart1(g *graph.Graph[Vertex], from, to Vertex) int {
	if !g.IsVertex(from) || !g.IsVertex(to) {
		panic(fmt.Sprintf("Internal error: [%s] or [%s] not a vertex", from, to))
	}

	vst := NewVisited(1)

	return countPaths(g, from, to, vst)
}

func CountPathsPart2(g *graph.Graph[Vertex], from, to Vertex) int {
	if !g.IsVertex(from) || !g.IsVertex(to) {
		panic(fmt.Sprintf("Internal error: [%s] or [%s] not a vertex", from, to))
	}

	vst := NewVisited(2)

	return countPaths(g, from, to, vst)
}

func countPaths(g *graph.Graph[Vertex], from, to Vertex, vst *Visited) int {
	if from == to {
		return 1
	}

	vst.Visit(from)

	neighbours := g.Neighbours(from)
	neighbours = fun.Filter(vst.CanVisit, neighbours)

	count := 0
	for _, neighbour := range neighbours {
		count += countPaths(g, neighbour, to, vst.Copy())
	}
	//	fmt.Printf("[%s -> %s]: neighbours [%s] returning %d\n", from, to, neighbours, count)
	return count
}

func lineToEdge(line string) graph.Edge[Vertex] {
	bits := strings.Split(line, "-")
	if len(bits) != 2 {
		panic(fmt.Sprintf("Bad edge-line [%s]", line))
	}
	return graph.Edge[Vertex]{From: Vertex(bits[0]), To: Vertex(bits[1])}
}

func NewDay12() *Day12 {
	d := Day12{}
	return &d
}

func (d *Day12) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")
	edges := fun.Map(lineToEdge, lines)
	fmt.Fprintf(out, "Edges: %v\n", edges)
	g := graph.NewFromEdges(edges, true)
	fmt.Fprintf(out, "G:\n%v\n", g)
	numPaths := CountPathsPart1(g, "start", "end")
	fmt.Fprintf(out, "Num paths: %d\n", numPaths)
	numPaths = CountPathsPart2(g, "start", "end")
	fmt.Fprintf(out, "Num paths: %d\n", numPaths)
	return nil
}
