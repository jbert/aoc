package y2025

import (
	"fmt"
	"io"

	"github.com/jbert/aoc"
	"github.com/jbert/aoc/num"
	"github.com/jbert/aoc/year"
	"github.com/jbert/fun"
)

type Day5 struct{ year.Year }

func (d *Day5) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")
	lgs := aoc.LineGroups(lines)
	if len(lgs) != 2 {
		panic("expected 2 line groups")
	}
	ranges, err := fun.ErrMap(parseRange, lgs[0])
	if err != nil {
		panic(fmt.Sprintf("can't parse range: %w"))
	}
	ids := fun.Map(num.MustAtoi, lgs[1])
	// fmt.Printf("ranges: %v\n", ranges)
	// fmt.Printf("ids: %v\n", ids)

	isFresh := func(n int) bool {
		for _, r := range ranges {
			if r.Contains(n) {
				return true
			}
		}
		return false
	}
	fmt.Printf("Part 1: %d\n", len(fun.Filter(isFresh, ids)))

	return nil
}
