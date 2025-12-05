package y2025

import (
	"fmt"
	"io"
	"slices"

	"github.com/jbert/aoc"
	"github.com/jbert/aoc/num"
	"github.com/jbert/aoc/year"
	"github.com/jbert/fun"
)

type Day5 struct{ year.Year }

func (r Range) Overlaps(s Range) bool {
	if r.lo > s.lo {
		return s.Overlaps(r)
	}
	return r.hi >= s.lo
}

func (r Range) Join(s Range) (Range, error) {
	if r.lo > s.lo {
		return s.Join(r)
	}

	if r.hi < s.lo {
		return Range{}, fmt.Errorf("%s, %s - no overlap", r, s)
	}
	if r.hi > s.hi {
		return r, nil
	}
	return Range{r.lo, s.hi}, nil
}

func (r Range) Size() int {
	return r.hi - r.lo + 1
}

func (d *Day5) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")
	lgs := aoc.LineGroups(lines)
	if len(lgs) != 2 {
		panic("expected 2 line groups")
	}
	ranges, err := fun.ErrMap(parseRange, lgs[0])
	if err != nil {
		panic(fmt.Sprintf("can't parse range: %s", err))
	}
	ids := fun.Map(num.MustAtoi, lgs[1])
	fmt.Printf("num ranges: %v\n", len(ranges))
	fmt.Printf("num ids: %v\n", len(ids))

	isFresh := func(n int) bool {
		for _, r := range ranges {
			if r.Contains(n) {
				return true
			}
		}
		return false
	}
	fmt.Printf("Part 1: %d\n", len(fun.Filter(isFresh, ids)))

	slices.SortFunc(ranges, cmpRange)

	var joinedRanges []Range
	for len(ranges) > 1 {
		r, err := ranges[0].Join(ranges[1])
		if err == nil {
			// fmt.Printf("JB0\n")
			ranges = ranges[1:]
			ranges[0] = r
		} else {
			// fmt.Printf("JB1\n")
			joinedRanges = append(joinedRanges, ranges[0])
			ranges = ranges[1:]
		}
	}
	joinedRanges = append(joinedRanges, ranges[0])
	fmt.Printf("%v\n", joinedRanges)
	sizes := fun.Map(func(r Range) int { return r.Size() }, joinedRanges)
	fmt.Printf("Part 2 : %d\n", fun.Sum(sizes))

	/*
		for i := range joinedRanges {
			for j := range joinedRanges {
				if j <= i {
					continue
				}
				if joinedRanges[i].Overlaps(joinedRanges[j]) {
					fmt.Printf("LOGIC ERROR: %d %d", i, j)
				}
				_, err := joinedRanges[i].Join(joinedRanges[j])
				if err == nil {
					fmt.Printf("LOGIC ERROR2: %d %d", i, j)
				}
			}
		}
	*/

	return nil
}
