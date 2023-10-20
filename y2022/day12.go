package y2022

import (
	"fmt"
	"io"
	"math"
	"sort"

	"github.com/jbert/aoc"
	"github.com/jbert/aoc/astar"
	"github.com/jbert/fun"
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

	start := findChar(g, 'S')
	goal := findChar(g, 'E')
	hc := func(p pts.P2) float64 { return float64(p.ManhattanLength()) }

	hg := HeightGrid{start: start, goal: goal, g: g}
	path, err := astar.Astar(start, goal, astar.Graph[pts.P2](hg), hc)
	if err != nil {
		return fmt.Errorf("Can't find path: %s", err)
	}
	fmt.Printf("Path: %v\n", path)
	fmt.Printf("Part 1: %d\n", len(path)-1)

	starts := findAll(g, 'a')
	starts = append(starts, findChar(g, 'S'))
	fmt.Printf("%d possible starts\n", len(starts))
	findPathLen := func(p pts.P2) int {
		path, err := astar.Astar(p, goal, astar.Graph[pts.P2](hg), hc)
		if err != nil {
			//			fmt.Printf("Can't find path: %s\n", err)
			return math.MaxInt
		}
		//		fmt.Printf("%s -> %d (%v)\n", p, len(path)-1, path)
		return len(path) - 1
	}
	lengths := fun.Map(findPathLen, starts)
	sort.Ints(lengths)
	fmt.Printf("Part 2: %d\n", lengths[0])
	return nil
}

func findAll(g grid.Grid[byte], c byte) []pts.P2 {
	var ps []pts.P2
	g.ForEachV(func(i, j int, v byte) {
		if v == c {
			ps = append(ps, pts.P2{i, j})
		}
	})
	return ps
}

func findChar(g grid.Grid[byte], c byte) pts.P2 {
	ps := findAll(g, c)
	if len(ps) != 1 {
		panic("wtf")
	}
	return ps[0]
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

	height := hg.g.GetPt(p)
	if height == 'S' {
		height = 'a'
	}
	canStep := func(q pts.P2) bool {
		qHeight := hg.g.GetPt(q)
		if qHeight == 'E' {
			qHeight = 'z'
		}
		if qHeight == 'S' {
			qHeight = 'a'
		}
		ok := qHeight < height || qHeight-height <= 1
		//		fmt.Printf("%s -> %s: %c -> %c (%c): %v (%d)\n", p, q, height, qHeight, hg.g.GetPt(q), ok, qHeight-height)
		return ok
	}
	filtered := fun.Filter(canStep, nps)
	//	fmt.Printf("%s: %c (NPS: %v)\n", p, height, filtered)
	return filtered
}

func (hg HeightGrid) Weight(from, to pts.P2) float64 {
	d := from.Sub(to)
	if d.ManhattanLength() != 1 {
		panic(fmt.Sprintf("wtf: from %s to %s", from, to))
	}
	return 1
}
