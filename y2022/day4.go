package y2022

import (
	"fmt"
	"io"
	"strings"

	"github.com/jbert/aoc/num"
	"github.com/jbert/fun"
)

type Day4 struct{}

func NewDay4() *Day4 {
	d := Day4{}
	return &d
}

func (d *Day4) Run(out io.Writer, lines []string) error {
	rangePairs := fun.Map(lineToRangePairs, lines)
	fmt.Printf("Part1: %d\n", len(fun.Filter(func(rp RangePair) bool {
		return rp.a.Covers(rp.b) || rp.b.Covers(rp.a)
	}, rangePairs)))
	fmt.Printf("Part1: %d\n", len(fun.Filter(func(rp RangePair) bool {
		return rp.a.Overlaps(rp.b)
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

func (r Range) Overlaps(s Range) bool {
	return !(r.lo > s.hi || r.hi < s.lo)
}

func (r Range) Covers(s Range) bool {
	return r.lo <= s.lo && r.hi >= s.hi
}

func parseRange(l string) Range {
	bits := strings.Split(l, "-")
	return Range{
		lo: num.MustAtoi(bits[0]),
		hi: num.MustAtoi(bits[1]),
	}
}

func lineToRangePairs(l string) RangePair {
	bits := strings.Split(l, ",")
	return RangePair{
		a: parseRange(bits[0]),
		b: parseRange(bits[1]),
	}
}
