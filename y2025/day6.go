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

func digToNum(dig byte) int {
	return int(dig - '0')
}

func joinDigits(digs []byte) int {
	x := 0
	digs = fun.Filter(func(b byte) bool { return b != ' ' }, digs)
	for i, dig := range digs {
		x += digToNum(dig)
		if i != len(digs)-1 {
			x *= 10
		}
	}
	return x
}

func cephParseBlock(lines []string, loCol int, hiCol int) []int {
	loCol += 1
	hiCol -= 1
	var xs []int
	for i := loCol; i <= hiCol; i++ {
		var digits []byte
		for j := range lines {
			digits = append(digits, lines[j][i])
		}
		x := joinDigits(digits)
		xs = append(xs, x)
	}
	return xs
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

	var spaceCols []int
COL:
	for i := range lines[0] {
		for j := range lines {
			if lines[j][i] != ' ' {
				continue COL
			}
		}
		spaceCols = append(spaceCols, i)
	}
	// fmt.Printf("%v\n", spaceCols)
	xss = [][]int{}

	spaceCols = append([]int{-1}, spaceCols...)
	spaceCols = append(spaceCols, len(lines[0]))
	for i := range spaceCols {
		if i == 0 {
			continue
		}
		xs := cephParseBlock(lines, spaceCols[i-1], spaceCols[i])
		xss = append(xss, xs)
	}

	// fmt.Printf("ops %v\n", ops)
	// fmt.Printf("xss %v\n", xss)
	sum = 0
	for i := range n {
		v := applyOp(ops[i], xss[i])
		sum += v
	}

	fmt.Printf("Part 2: %d\n", sum)

	return nil
}
