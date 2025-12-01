package y2021

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/jbert/aoc/astar"
)

type Day15 struct{}

func NewDay15() *Day15 {
	d := Day15{}
	return &d
}

func (d *Day15) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")
	grid := GridFromLines(lines)
	fmt.Printf("Grid:\n%v\n", grid)
	lowestRisk := grid.LowestRisk()
	fmt.Printf("Lowest risk: %d\n", lowestRisk)

	expandedGrid := expandGrid(grid)
	//	fmt.Printf("Grid:\n%v\n", expandedGrid)
	//	for y := range expandedGrid {
	//		fmt.Printf("%d", expandedGrid[y][expandedGrid.Width()-1])
	//	}
	//	fmt.Printf("\n")
	lowestRisk = expandedGrid.LowestRisk()
	fmt.Printf("Lowest risk: %d\n", lowestRisk)

	return nil
}

func expandGrid(grid Grid) Grid {
	eGrid := make([][]int, len(grid)*5)
	for y, row := range grid {
		for j := 0; j < 5; j++ {
			eGrid[y+j*grid.Height()] = make([]int, grid.Width()*5)
		}
		for x, risk := range row {
			for i := 0; i < 5; i++ {
				ex := x + grid.Width()*i
				for j := 0; j < 5; j++ {
					ey := y + grid.Height()*j
					eRisk := risk + i + j
					for eRisk > 9 {
						eRisk -= 9
					}
					eGrid[ey][ex] = eRisk
				}
			}
		}
	}
	return eGrid
}

type Grid [][]int

func (g Grid) Width() int {
	return len(g[0])
}

func (g Grid) Height() int {
	return len(g)
}

func GridFromLines(lines []string) Grid {
	var grid [][]int
	for _, line := range lines {
		var row []int
		digits := strings.Split(line, "")
		for _, digit := range digits {
			num, err := strconv.Atoi(digit)
			if err != nil {
				panic(fmt.Sprintf("Bad digit? [%s]: %s", digit, err))
			}
			row = append(row, num)
		}
		grid = append(grid, row)
	}

	return grid
}

func (grid Grid) LowestRisk() int {
	start := Pos{0, 0}
	goal := Pos{grid.Width() - 1, grid.Height() - 1}
	graph := astar.Graph[Pos](grid)
	riskPath, err := astar.Astar(start, goal, graph, func(p Pos) float64 { return 0 })
	if err != nil {
		panic("Couldn't find path")
	}
	return grid.PathRisk(riskPath)
}

// Satisfy github.com/jbert/aoc2021/astar interface
type Pos struct {
	X int
	Y int
}

func (g Grid) Neighbours(p Pos) []Pos {
	var ns []Pos

	if p.X > 0 {
		ns = append(ns, Pos{p.X - 1, p.Y})
	}
	if p.Y > 0 {
		ns = append(ns, Pos{p.X, p.Y - 1})
	}
	if p.X < g.Width()-1 {
		ns = append(ns, Pos{p.X + 1, p.Y})
	}
	if p.Y < g.Height()-1 {
		ns = append(ns, Pos{p.X, p.Y + 1})
	}
	return ns
}

func (g Grid) Weight(from, to Pos) float64 {
	return float64(g.Lookup(to))
}

func (g Grid) Lookup(p Pos) int {
	return g[p.Y][p.X]
}

func (g Grid) PathRisk(path []Pos) int {
	risk := 0
	// Don't count first
	first := true
	for _, p := range path {
		if !first {
			risk += g.Lookup(p)
		} else {
			first = false
		}
	}
	return risk
}
