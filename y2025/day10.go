package y2025

import (
	"fmt"
	"io"
	"strings"

	"github.com/jbert/aoc"
	"github.com/jbert/aoc/year"
	"github.com/jbert/fun"
)

type Day10 struct{ year.Year }

type machine struct {
	wanted   int
	wiring   [][]int
	joltages []int
}

func (m machine) String() string {
	return fmt.Sprintf("%b: %v {%v}", m.wanted, m.wiring, m.joltages)
}

func parseWanted(s string) int {
	s = s[1 : len(s)-1]
	bs := []byte(s)
	// bs = fun.Reverse(bs)
	n := 0
	for _, b := range bs {
		n <<= 1
		if b == '#' {
			n += 1
		}
	}
	return n
}

func parseWiring(s string) []int {
	s = s[1 : len(s)-1]
	return aoc.StringToInts(s)
}

func machineFromString(s string) *machine {
	bits := strings.Split(s, " ")
	wanted := parseWanted(bits[0])
	wiring := fun.Map(parseWiring, bits[1:len(bits)-1])
	jStr := bits[len(bits)-1]
	joltages := aoc.StringToInts(jStr[1 : len(jStr)-1])
	return &machine{
		wanted:   wanted,
		wiring:   wiring,
		joltages: joltages,
	}
}

func (d *Day10) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")
	ms := fun.Map(machineFromString, lines)
	for _, m := range ms {
		fmt.Printf("%s\n", m)
	}

	// fmt.Printf("Part 1: %d\n", fun.Sum(joltages))
	// fmt.Printf("Part 2: %d\n", fun.Sum(joltages))

	return nil
}
