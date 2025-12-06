package y2025

import (
	"fmt"
	"io"
	"strings"

	"github.com/jbert/aoc/num"
	"github.com/jbert/aoc/year"
	"github.com/jbert/fun"
)

type Day6 struct{ year.Year }

func lineToBits(l string) []string {
	bits := strings.Split(l, " ")
	bits = fun.Filter(func(s string) bool { return s != "" }, bits)
	return bits
}

func applyOp(op string, xs []int) int {
	switch op {
	case "+":
		return fun.Sum(xs)
	case "*":
		return fun.Prod(xs)
	default:
		panic(fmt.Sprintf("unknown op: %s", op))
	}
}

func (d *Day6) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")

	ops := lineToBits(lines[len(lines)-1])
	n := len(ops)
	lines = lines[:len(lines)-1]

	var xss [][]int
	for _, l := range lines {
		xs := fun.Map(num.MustAtoi, lineToBits(l))
		if len(xs) != n {
			return fmt.Errorf("diff number of ops (%d) and xs (%d)", n, len(xs))
		}
		xss = append(xss, xs)
	}
	// fmt.Printf("%v\n", ops)
	// fmt.Printf("%v\n", xss)
	sum := 0
	for i := range n {
		var xs []int
		for _, l := range xss {
			xs = append(xs, l[i])
		}
		v := applyOp(ops[i], xs)
		sum += v
	}

	fmt.Printf("Part 1: %d\n", sum)
	// fmt.Printf("Part 2: %d\n", fun.Sum(joltages))

	return nil
}
