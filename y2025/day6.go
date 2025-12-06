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

// zero indexed n'th digit
// 123, 0 -> 1
// 123, 2 -> 3
func nthDigit(x int, j int) int {
	// i := numDigits(x) - j - 1
	factor := pow10(j + 1)
	// fmt.Printf("ND: x %d j %d i %d factor %d\n", x, j, i, factor)
	x /= factor
	// fmt.Printf("ND:  %d j %d -> %d\n", x*factor, j, x%10)
	return x % 10
}

func joinDigits(digs []int) int {
	x := 0
	for i, dig := range digs {
		x += dig
		if i != len(digs)-1 {
			x *= 10
		}
	}
	return x
}

func cephTransform(xs []int) []int {
	numCols := 0
	for _, x := range xs {
		numDig := numDigits(x)
		if numDig > numCols {
			numCols = numDig
		}
	}

	offsets := make([]int, len(xs))
	for i, x := range xs {
		numDig := numDigits(x)
		offsets[i] = numCols - numDig
	}

	var cxs []int
	for i := range numCols {

		var cxDig []int
		for j, x := range xs {
			if i >= offsets[j] {
				cxDig = append(cxDig, nthDigit(x, i-offsets[j]))
			}
		}
		cx := joinDigits(cxDig)
		cxs = append(cxs, cx)
	}
	return cxs
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

	sum = 0
	for i := range n {
		var xs []int
		for _, l := range xss {
			xs = append(xs, l[i])
		}
		xs = cephTransform(xs)
		// fmt.Printf("xs %d\n", xs)
		v := applyOp(ops[i], xs)
		sum += v
	}

	fmt.Printf("Part 2: %d\n", sum)

	return nil
}
