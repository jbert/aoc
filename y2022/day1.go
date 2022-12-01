package y2022

import (
	"fmt"
	"io"

	"github.com/jbert/aoc"
	"github.com/jbert/aoc/fun"
)

type Day1 struct{ Year }

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
		ns := fun.Map(aoc.MustAtoi, lg)
		calories := fun.Sum(ns)
		if calories > maxCalories {
			maxCalories = calories
		}
	}
	fmt.Printf("Part1: %d\n", maxCalories)
	return nil
}
