package y2022

import (
	"fmt"
	"io"
	"strings"

	"github.com/jbert/aoc/fun"
	"github.com/jbert/aoc/graph"
	"github.com/jbert/aoc/num"
)

type Day16 struct{ Year }

func NewDay16() *Day16 {
	d := Day16{}
	return &d
}

type Valve struct {
	label    string
	flowRate int
	on       bool
}

func (d *Day16) Run(out io.Writer, lines []string) error {
	valves := make(map[string]Valve)
	var edges []graph.Edge[string]
	for _, l := range lines {
		v, lineEdges := parseLine(l)
		valves[v.label] = v
		edges = append(edges, lineEdges...)
	}
	g := graph.NewFromEdges(edges, true)
	fmt.Printf("G: %v\n", g)
	fmt.Printf("V: %v\n", valves)
	return nil
}

func parseLine(l string) (Valve, []graph.Edge[string]) {
	v := Valve{}
	v.label = l[6:8]
	l = l[23:]
	i := strings.Index(l, ";")
	v.flowRate = num.MustAtoi(l[:i])
	l = l[i:]
	l = l[18:]
	i = strings.Index(l, " ")
	l = l[i+1:]
	fmt.Printf("[%s]\n", l)
	dests := strings.Split(l, ", ")
	edges := fun.Map(func(s string) graph.Edge[string] { return graph.Edge[string]{From: v.label, To: s} }, dests)
	return v, edges
}
