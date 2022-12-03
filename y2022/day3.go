package y2022

import (
	"fmt"
	"io"

	"github.com/jbert/aoc/fun"
	"github.com/jbert/aoc/set"
)

type Day3 struct{ Year }

func NewDay3() *Day3 {
	d := Day3{}
	return &d
}

func repeatedType(s string) byte {
	buf := []byte(s)
	if len(buf)%2 != 0 {
		panic("odd length string")
	}
	b2 := len(buf) / 2
	a := set.SetFromList(buf[:b2])
	b := set.SetFromList(buf[b2:])
	return a.Intersect(b).ToList()[0]
}

func typeToPrio(t byte) int {
	if t >= 'a' && t <= 'z' {
		return int(t - 'a' + 1)
	}
	if t >= 'A' && t <= 'Z' {
		return int(t - 'A' + 27)
	}
	panic("bad type")
}

func (d *Day3) Run(out io.Writer, lines []string) error {

	types := fun.Map(repeatedType, lines)
	fmt.Printf("types: %v\n", types)
	prios := fun.Map(typeToPrio, types)
	fmt.Printf("prios: %v\n", prios)
	fmt.Printf("Part1: %d\n", fun.Sum(prios))

	return nil
}
