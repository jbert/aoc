package y2022

import (
	"fmt"
	"io"
	"strings"

	"github.com/jbert/aoc"
	"github.com/jbert/fun"
	"github.com/jbert/aoc/num"
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
	// fmt.Printf("Stacks: %v\n", stacks)
	printStacks(stacks)
	moves := fun.Map(moveFromLine, lgs[1])
	fmt.Printf("Moves:\n%v\n", moves)
	applyMove := func(m StackMove, stacks []stack.Stack[byte]) []stack.Stack[byte] {
		m.Apply(stacks)
		fmt.Printf("\nM: %s\n", m)
		printStacks(stacks)
		return stacks
	}
	fun.Foldl(applyMove, stacks, moves)
	fmt.Printf("Part1: ")
	for _, s := range stacks {
		fmt.Printf("%c", s.MustPeek())
	}
	fmt.Printf("\n")

	// Reset for part 2
	stacks = linesToStacks(lgs[0])
	applyMoveP2 := func(m StackMove, stacks []stack.Stack[byte]) []stack.Stack[byte] {
		m.ApplyP2(stacks)
		fmt.Printf("\nM: %s\n", m)
		printStacks(stacks)
		return stacks
	}
	fun.Foldl(applyMoveP2, stacks, moves)
	fmt.Printf("Part2: ")
	for _, s := range stacks {
		fmt.Printf("%c", s.MustPeek())
	}
	fmt.Printf("\n")

	return nil
}

func printStacks(stacks []stack.Stack[byte]) {
	for i, s := range stacks {
		fmt.Printf("%d ", i+1)
		for _, c := range s {
			fmt.Printf("%c", c)
		}
		fmt.Printf("\n")
	}
}

type StackMove struct {
	from, to, amount int
}

func (m StackMove) String() string {
	return fmt.Sprintf("%d from %d -> %d", m.amount, m.from, m.to)
}

func (m StackMove) ApplyP2(stacks []stack.Stack[byte]) {
	var moved []byte
	for i := 0; i < m.amount; i++ {
		x := stacks[m.from-1].MustPop()
		moved = append(moved, x)
	}
	moved = fun.Reverse(moved)
	for _, x := range moved {
		stacks[m.to-1].Push(x)
	}
}

func (m StackMove) Apply(stacks []stack.Stack[byte]) {
	for i := 0; i < m.amount; i++ {
		x := stacks[m.from-1].MustPop()
		stacks[m.to-1].Push(x)
	}
}

func moveFromLine(l string) StackMove {
	bits := strings.Split(l, " ")
	return StackMove{
		from:   num.MustAtoi(bits[3]),
		to:     num.MustAtoi(bits[5]),
		amount: num.MustAtoi(bits[1]),
	}
}

func linesToStacks(lines []string) []stack.Stack[byte] {
	// Drop labelling line
	lines = lines[:len(lines)-1]

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
