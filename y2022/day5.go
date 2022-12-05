package y2022

import (
	"fmt"
	"io"

	"github.com/jbert/aoc"
	"github.com/jbert/aoc/fun"
	"github.com/jbert/aoc/stack"
)

type Day5 struct{ Year }

func NewDay5() *Day5 {
	d := Day5{}
	return &d
}

func (d *Day5) Run(out io.Writer, lines []string) error {
	lgs := aoc.LineGroups(lines)
	fmt.Printf("%s\n", lgs[0])
	fmt.Printf("%s\n", lgs[1])
	stacks := linesToStacks(lgs[0])
	fmt.Printf("Stacks: %v\n", stacks)
	return nil
}

func linesToStacks(lines []string) []stack.Stack[byte] {
	// Drop labelling line
	lines = lines[:len(lines)-1]
	lines = fun.Reverse(lines)

	numStacks := (len(lines[0]) + 1) / 4
	stacks := make([]stack.Stack[byte], numStacks)
	for i := range stacks {
		stacks[i] = stack.New[byte]()
	}
	for _, l := range fun.Reverse(lines) {
		for j := 0; j < (len(l)+1)/4; j++ {
			c := l[j*4+1]
			if c != ' ' {
				stacks[j].Push(l[j*4+1])
			}
		}
	}
	return stacks
}
