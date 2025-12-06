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

func applyOp(op string, a int, b int, c int) int {
	switch op {
	case "+":
		return a + b + c
	case "*":
		return a * b * c
	default:
		panic(fmt.Sprintf("unknown op: %s", op))
	}
}

func (d *Day6) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")
	as := fun.Map(num.MustAtoi, lineToBits(lines[0]))
	bs := fun.Map(num.MustAtoi, lineToBits(lines[1]))
	cs := fun.Map(num.MustAtoi, lineToBits(lines[2]))
	ops := lineToBits(lines[3])
	fmt.Printf("%v\n", as)
	fmt.Printf("%v\n", bs)
	fmt.Printf("%v\n", cs)
	fmt.Printf("%v\n", ops)
	n := len(as)
	if len(bs) != n {
		return fmt.Errorf("diff number of as (%d) and bs (%d)", n, len(bs))
	}
	if len(cs) != n {
		return fmt.Errorf("diff number of as (%d) and cs (%d)", n, len(cs))
	}
	if len(ops) != n {
		return fmt.Errorf("diff number of as (%d) and ops(%d)", n, len(ops))
	}
	sum := 0
	for i := range n {
		v := applyOp(ops[i], as[i], bs[i], cs[i])
		sum += v
	}

	fmt.Printf("Part 1: %d\n", sum)
	// fmt.Printf("Part 2: %d\n", fun.Sum(joltages))

	return nil
}
