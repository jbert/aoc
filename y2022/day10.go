package y2022

import (
	"fmt"
	"io"
	"strings"

	"github.com/jbert/aoc"
	"github.com/jbert/aoc/fun"
)

type Day10 struct{ Year }

func NewDay10() *Day10 {
	d := Day10{}
	return &d
}

func (d *Day10) Run(out io.Writer, lines []string) error {
	instructions := fun.Map(lineToInstruction, lines)
	cpu := NewCPU()
	// Keep cpu fed with instructions
	readAt := []int{20, 60, 100, 140, 180, 220}
	readings := []int{}
	cpu.monitor = func(x int, cycle int) {
		if len(readAt) == 0 {
			return
		}
		if cycle == readAt[0] {
			readings = append(readings, cycle*x)
			fmt.Printf("%d: %d * %d == %d\n", cycle, cycle, x, cycle*x)
			readAt = readAt[1:]
		}
	}
	for _, inst := range instructions {
		cpu.Execute(inst)
	}
	fmt.Printf("readings: %v\n", readings)
	fmt.Printf("Part 1: %d\n", fun.Sum(readings))
	return nil
}

func lineToInstruction(l string) Instruction {
	bits := strings.Split(l, " ")
	if len(bits) == 1 {
		return Instruction{kind: bits[0]}
	}
	return Instruction{kind: bits[0], arg: aoc.MustAtoi(bits[1])}
}

type Instruction struct {
	kind string
	arg  int
}

func (i Instruction) apply(c *CPU) {
	switch i.kind {
	case "noop":
		return
	case "addx":
		c.x += i.arg
	default:
		panic("wtf")
	}
}

func (i Instruction) cycles() int {
	switch i.kind {
	case "noop":
		return 1
	case "addx":
		return 2
	default:
		panic("wtf")
	}
}

type CPU struct {
	x       int
	cycle   int
	monitor func(x int, cycle int)
}

func NewCPU() *CPU {
	return &CPU{x: 1, cycle: 0}
}

func (c *CPU) Execute(inst Instruction) {
	cycles := inst.cycles()
	for i := 0; i < cycles; i++ {
		c.cycle++
		c.monitor(c.x, c.cycle)
	}
	inst.apply(c)
}
