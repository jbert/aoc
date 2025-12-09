package y2025

import (
	"fmt"
	"io"
	"slices"

	"github.com/jbert/aoc/pts"
	"github.com/jbert/aoc/year"
	"github.com/jbert/fun"
)

type Day9 struct{ year.Year }

func (d *Day9) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")
	ps := fun.Map(pts.P2FromString, lines)
	fmt.Printf("%v\n", ps)
	slices.SortFunc(ps, pts.CmpP2)

	mxArea := 0
	// var mxRect pts.Rect
	for i := range ps {
		for j := range ps {
			if j <= i {
				continue
			}
			r := pts.NewRect(ps[i], ps[j])
			ra := r.Area()
			// fmt.Printf("RA %d R: %s\n", ra, r)
			if ra > mxArea {
				mxArea = ra
				// mxRect = r
			}
		}
	}
	// fmt.Printf("mx %d mxRect: %v\n", mxArea, mxRect)
	fmt.Printf("Part 1: %d\n", mxArea)
	// fmt.Printf("Part 2: %d\n", fun.Sum(joltages))

	return nil
}
