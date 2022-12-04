package y2022

import (
	"fmt"
	"io"
	"strings"

	"github.com/jbert/aoc"
	"github.com/jbert/aoc/fun"
)

type Day4 struct{ Year }

func NewDay4() *Day4 {
	d := Day4{}
	return &d
}

func (d *Day4) Run(out io.Writer, lines []string) error {
	rangePairs := fun.Map(lineToRangePairs, lines)
	fmt.Printf("Part1: %d\n", len(fun.Filter(func(rp RangePair) bool {
		return rp.a.Covers(rp.b) || rp.b.Covers(rp.a)
	}, rangePairs)))
	return nil
}

type RangePair struct {
	a, b Range
}

type Range struct {
	lo int
	hi int
}

func (r Range) Covers(s Range) bool {
	return r.lo <= s.lo && r.hi >= s.hi
}

func parseRange(l string) Range {
	bits := strings.Split(l, "-")
	return Range{
		lo: aoc.MustAtoi(bits[0]),
		hi: aoc.MustAtoi(bits[1]),
	}
}

func lineToRangePairs(l string) RangePair {
	bits := strings.Split(l, ",")
	return RangePair{
		a: parseRange(bits[0]),
		b: parseRange(bits[1]),
	}
}
