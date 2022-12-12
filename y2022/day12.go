package y2022

import (
	"fmt"
	"io"

	"github.com/jbert/aoc"
	"github.com/jbert/aoc/astar"
	"github.com/jbert/aoc/fun"
	"github.com/jbert/aoc/grid"
	"github.com/jbert/aoc/pts"
)

type Day12 struct{ Year }

func NewDay12() *Day12 {
	d := Day12{}
	return &d
}

func (d *Day12) Run(out io.Writer, lines []string) error {
	g := aoc.ByteGrid(lines)
	fmt.Printf("%s\n", g)

	start := pts.P2{0, 0}
	goal := findGoal(g)
	hc := func(p pts.P2) float64 { return 1.0 }

	hg := HeightGrid{start: start, goal: goal, g: g}
	path, err := astar.Astar(start, goal, astar.Graph[pts.P2](hg), hc)
	if err != nil {
		return fmt.Errorf("Can't find path: %s", err)
	}
	fmt.Printf("Path: %v\n", path)
	fmt.Printf("Part 1: %d\n", len(path)-1)
	return nil
}

func findGoal(g grid.Grid[byte]) pts.P2 {
	found := false
	var p pts.P2
	g.ForEachV(func(i, j int, v byte) {
		if v == 'E' {
			found = true
			p = pts.P2{i, j}
		}
	})
	if !found {
		panic("wtf")
	}
	return p
}

// Implement astar.Graph interface over the byte grid and puzzle rules
/*
type Graph[V comparable] interface {
	Neighbours(V) []V
	Weight(from, to V) float64
}
*/

type HeightGrid struct {
	start pts.P2
	goal  pts.P2
	g     grid.Grid[byte]
}

func (hg HeightGrid) Neighbours(p pts.P2) []pts.P2 {
	nps := hg.g.CardinalNeighbourPts(p)
	//	fmt.Printf("NPS: %v (%v)\n", nps, p)

	height := hg.g.GetPt(p)
	canStep := func(q pts.P2) bool {
		return p.Equals(hg.start) || (height == 'z' && q.Equals(hg.goal)) || hg.g.GetPt(q)-height <= 1
	}
	filtered := fun.Filter(canStep, nps)
	return filtered
}

func (hg HeightGrid) Weight(from, to pts.P2) float64 {
	d := from.Sub(to)
	if d.ManhattanLength() != 1 {
		panic(fmt.Sprintf("wtf: from %s to %s", from, to))
	}
	return 1
}
