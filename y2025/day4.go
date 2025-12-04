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

func (d *Day4) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")
	rows := fun.Map(func(s string) []string { return strings.Split(s, "") }, lines)

	g := grid.NewFromRows(rows)

	hasPaper := func(p pts.P2) bool { return g.GetPt(p) == "@" }
	countNeighbours := func(p pts.P2) int {
		nsWithPaper := fun.Filter(hasPaper, g.AllNeighbourPts(p))
		// fmt.Printf("p %v nsWithPaper %v\n", p, nsWithPaper)
		// fmt.Printf("anp %v\n", g.AllNeighbourPts(p))

		return len(nsWithPaper)
	}

	numNs := grid.NewFromFunc(g.Width(), g.Height(), countNeighbours)
	for _, row := range numNs {
		fmt.Printf("%v\n", row)
	}
	count := 0
	numNs.ForEach(func(p pts.P2) {
		if hasPaper(p) && numNs.GetPt(p) < 4 {
			count += 1
		}
	})
	fmt.Printf("Part 1: %d\n", count)

	return nil
}
