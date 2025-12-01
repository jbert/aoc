package y2022

import (
	"fmt"
	"io"
	"sort"

	"github.com/jbert/aoc"
	"github.com/jbert/aoc/num"
	"github.com/jbert/fun"
)

type Day1 struct{}

func NewDay1() *Day1 {
	d := Day1{}
	return &d
}

func (d *Day1) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")
	lgs := aoc.LineGroups(lines)
	fmt.Printf("LGS:\n%v\n", lgs)

	maxCalories := 0
	for _, lg := range lgs {
		ns := fun.Map(num.MustAtoi, lg)
		calories := fun.Sum(ns)
		if calories > maxCalories {
			maxCalories = calories
		}
	}
	fmt.Printf("Part1: %d\n", maxCalories)

	ngs := fun.Map(func(lg []string) []int {
		return fun.Map(num.MustAtoi, lg)
	}, lgs)
	// Sort biggest first
	sort.Slice(ngs, func(i, j int) bool {
		return fun.Sum(ngs[i]) > fun.Sum(ngs[j])
	})
	sums := fun.Map(fun.Sum[int], ngs)
	fmt.Printf("SUMS: %v\n", sums)
	fmt.Printf("Part2: %d\n", fun.Sum(sums[:3]))

	return nil
}
