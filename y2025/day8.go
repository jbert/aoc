package y2025

import (
	"fmt"
	"io"
	"strings"

	"github.com/jbert/aoc/num"
	"github.com/jbert/aoc/pts"
	"github.com/jbert/aoc/year"
	"github.com/jbert/fun"
)

type Day8 struct{ year.Year }

func parsePoint(l string) pts.P3 {
	bits := strings.Split(l, ",")
	ns := fun.Map(num.MustAtoi, bits)
	return pts.P3{ns[0], ns[1], ns[2]}
}

func (d *Day8) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")
	pts := fun.Map(parsePoint, lines)
	fmt.Printf("pts %v\n", pts)
	// fmt.Printf("Part 1: %d\n", fun.Sum(joltages))
	// fmt.Printf("Part 2: %d\n", fun.Sum(joltages))

	return nil
}
