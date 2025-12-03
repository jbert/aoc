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

func maxJoltage(b []int) int {
	l := len(b)
	ifirst, first := findMax(b[:l-1])
	_, second := findMax(b[ifirst+1:])
	return first*10 + second
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
	joltages := fun.Map(maxJoltage, banks)
	fmt.Printf("Part 1: %d\n", fun.Sum(joltages))

	return nil
}
