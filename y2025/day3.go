package y2025

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/jbert/aoc/year"
	"github.com/jbert/fun"
)

type Day3 struct{ year.Year }

func findMax(b []int) (int, int) {
	imx := -1
	mx := -1
	for i, v := range b {
		if v > mx {
			mx = v
			imx = i
		}
	}
	if mx == -1 {
		panic("no max")
	}
	return imx, mx
}

func maxJoltage2(b []int) int {
	// l := len(b)
	// ifirst, first := findMax(b[:l-1])
	// _, second := findMax(b[ifirst+1:])
	// return first*10 + second
	return maxJoltageN(b, 2)
}

func maxJoltageN(b []int, n int) int {
	if n == 0 {
		return 0
	}
	l := len(b)
	i, v := findMax(b[:l-(n-1)])
	rest := maxJoltageN(b[i+1:], n-1)
	return v*pow10(n) + rest
}

func mustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("mustAtoi failed on [%s]: %s", s, err))
	}
	return n
}

func parseBank(l string) []int {
	digits := strings.Split(l, "")
	nums := fun.Map(mustAtoi, digits)
	return nums
}

func (d *Day3) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")
	banks := fun.Map(parseBank, lines)
	// fmt.Printf("banks: %+v\n", banks)

	// for _, b := range banks {
	// fmt.Fprintf(out, "%d\n", maxJoltage(b))
	// }
	joltages := fun.Map(maxJoltage2, banks)
	fmt.Printf("Part 1: %d\n", fun.Sum(joltages))
	joltages = fun.Map(func(b []int) int { return maxJoltageN(b, 12) }, banks)
	fmt.Printf("Part 2: %d\n", fun.Sum(joltages))

	return nil
}
