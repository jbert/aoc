package y2022

import (
	"fmt"
	"io"

	"github.com/jbert/fun"
	"github.com/jbert/set"
)

type Day3 struct{}

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
	a := set.NewFromList(buf[:b2])
	b := set.NewFromList(buf[b2:])
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

func elfGroupToBadge(lines []string) byte {
	a := set.NewFromList([]byte(lines[0]))
	b := set.NewFromList([]byte(lines[1]))
	c := set.NewFromList([]byte(lines[2]))

	return a.Intersect(b).Intersect(c).ToList()[0]
}

func (d *Day3) Run(out io.Writer, lines []string) error {

	types := fun.Map(repeatedType, lines)
	fmt.Printf("types: %v\n", types)
	prios := fun.Map(typeToPrio, types)
	fmt.Printf("prios: %v\n", prios)
	fmt.Printf("Part1: %d\n", fun.Sum(prios))

	elfGroups := fun.SplitBy(lines, 3)
	badges := fun.Map(elfGroupToBadge, elfGroups)
	fmt.Printf("Part2: %d\n", fun.Sum(fun.Map(typeToPrio, badges)))

	return nil
}
