package y2022

import (
	"fmt"
	"io"
	"strings"

	"github.com/jbert/aoc/grid"
	"github.com/jbert/aoc/num"
	"github.com/jbert/aoc/pts"
)

type Day14 struct{}

func NewDay14() *Day14 {
	d := Day14{}
	return &d
}

func (d *Day14) Run(out io.Writer, lines []string) error {
	g := linesToGrid(lines)
	fmt.Printf("%s\n", gridString(g))

	sandSource := pts.P2{500, 0}
	units := 0
	for {
		if !addSand(g, sandSource) {
			break
		}
		units++
	}
	fmt.Printf("%s\n", gridString(g))
	fmt.Printf("Part 1: %d\n", units)

	g = linesToGrid(lines)
	floorY := g.MaxY + 2
	floorSafeMinX := sandSource.X - 2*floorY
	floorSafeMaxX := sandSource.X + 2*floorY
	addLine(g, pts.P2{floorSafeMinX, floorY}, pts.P2{floorSafeMaxX, floorY}, '#')

	units = 0
	for g.GetPt(sandSource) == '.' {
		addSand(g, sandSource)
		units++
	}
	fmt.Printf("%s\n", gridString(g))
	fmt.Printf("Part 2: %d\n", units)

	return nil
}

func linesToGrid(lines []string) *grid.Sparse[byte] {
	g := grid.NewSparse[byte]('.')
	for _, l := range lines {
		bits := strings.Split(l, " ")
		current := pts.P2FromString(bits[0])
		for _, b := range bits[1:] {
			if b == "->" {
				continue
			}
			next := pts.P2FromString(b)
			addLine(g, current, next, '#')
			current = next
		}
	}
	return g
}

func addSand(g *grid.Sparse[byte], sand pts.P2) bool {
	down := pts.P2{0, +1}
	downLeft := pts.P2{-1, +1}
	downRight := pts.P2{+1, +1}
	tries := []pts.P2{down, downLeft, downRight}

STEP:
	for sand.Y <= g.MaxY {
		for _, try := range tries {
			possible := sand.Add(try)
			if g.GetPt(possible) == '.' {
				sand = possible
				continue STEP
			}
		}
		g.SetPt(sand, 'o')
		return true
	}
	return false
}

func gridString(g *grid.Sparse[byte]) string {
	b := &strings.Builder{}
	fmt.Printf("X: %d - %d, Y %d - %d\n\n", g.MinX, g.MaxX, g.MinY, g.MaxY)
	for y := g.MinY; y <= g.MaxY; y++ {
		for x := g.MinX; x <= g.MaxX; x++ {
			fmt.Fprintf(b, "%c", g.GetPt(pts.P2{x, y}))
		}
		fmt.Fprintf(b, "\n")
	}
	return b.String()
}

func addLine(g *grid.Sparse[byte], from pts.P2, to pts.P2, c byte) {
	//	fmt.Printf("Line from [%s] to [%s]\n", from, to)
	d := to.Sub(from)
	if d.X != 0 && d.Y != 0 {
		panic(fmt.Sprintf("Line not recti: %s", d))
	}
	d.X = num.Sign(d.X)
	d.Y = num.Sign(d.Y)
	for p := from; !p.Equals(to); p = p.Add(d) {
		g.SetPt(p, c)
	}
	g.SetPt(to, c)
}
