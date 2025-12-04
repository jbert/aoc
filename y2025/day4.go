package y2025

import (
	"fmt"
	"io"
	"strings"

	"github.com/jbert/aoc/grid"
	"github.com/jbert/aoc/pts"
	"github.com/jbert/aoc/year"
	"github.com/jbert/fun"
)

type Day4 struct{ year.Year }

func removablePts(g grid.Grid[string]) []pts.P2 {
	hasPaper := func(p pts.P2) bool { return g.GetPt(p) == "@" }
	countNeighbours := func(p pts.P2) int {
		nsWithPaper := fun.Filter(hasPaper, g.AllNeighbourPts(p))
		// fmt.Printf("p %v nsWithPaper %v\n", p, nsWithPaper)
		// fmt.Printf("anp %v\n", g.AllNeighbourPts(p))

		return len(nsWithPaper)
	}

	numNs := grid.NewFromFunc(g.Width(), g.Height(), countNeighbours)
	// for _, row := range numNs {
	// fmt.Printf("%v\n", row)
	// }
	var ps []pts.P2
	numNs.ForEach(func(p pts.P2) {
		if hasPaper(p) && numNs.GetPt(p) < 4 {
			ps = append(ps, p)
		}
	})
	return ps
}

func (d *Day4) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")
	rows := fun.Map(func(s string) []string { return strings.Split(s, "") }, lines)

	g := grid.NewFromRows(rows)

	remPts := removablePts(g)
	fmt.Printf("Part 1: %d\n", len(remPts))

	count := 0
	for len(remPts) > 0 {
		p := remPts[0]
		g.SetPt(p, ".")
		count += 1
		remPts = removablePts(g)
	}
	fmt.Printf("Part 2: %d\n", count)

	return nil
}
