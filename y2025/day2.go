package y2025

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/jbert/aoc/year"
	"github.com/jbert/fun"
)

type Day2 struct{ year.Year }

type Range struct {
	lo int
	hi int
}

func (r Range) String() string {
	return fmt.Sprintf("%d->%d", r.lo, r.hi)
}

func parseRange(s string) (*Range, error) {
	bits := strings.Split(s, "-")
	if len(bits) != 2 {
		return nil, fmt.Errorf("can't parse [%s]: not two bits", s)
	}
	lo, err := strconv.Atoi(bits[0])
	if err != nil {
		return nil, fmt.Errorf("can't parse [%s] as int: %w", bits[0], err)
	}
	hi, err := strconv.Atoi(bits[1])
	if err != nil {
		return nil, fmt.Errorf("can't parse [%s] as int: %w", bits[0], err)
	}
	return &Range{lo, hi}, nil
}

func (d *Day2) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")
	fmt.Fprintf(out, "%+v\n", lines)
	rangeStrs := strings.Split(lines[0], ",")
	fmt.Fprintf(out, "%+v\n", rangeStrs)
	ranges, err := fun.ErrMap(parseRange, rangeStrs)
	if err != nil {
		return fmt.Errorf("can't parse: %w", err)
	}
	fmt.Fprintf(out, "%+v\n", ranges)

	return nil
}
