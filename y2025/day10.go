package y2025

import (
	"fmt"
	"io"
	"strings"

	"github.com/jbert/aoc"
	"github.com/jbert/aoc/year"
	"github.com/jbert/fun"
	"github.com/jbert/set"
)

type Day10 struct{ year.Year }

type machine struct {
	wanted    int
	wiring    []int
	buttons   []set.Set[int]
	joltages  []int
	numLights int
}

func (m machine) numJoltages() int {
	// Part 1 and Part 2 have different goals, but same number
	return m.numLights
}

func (m machine) String() string {
	return fmt.Sprintf("%0*b: %v {%v} [%d]", m.numLights, m.wanted, m.wiring, m.joltages, m.numLights)
}

func (m machine) numButtons() int {
	return len(m.wiring)
}

func parseWanted(s string) int {
	s = s[1 : len(s)-1]
	bs := []byte(s)
	bs = fun.Reverse(bs)
	n := 0
	for _, b := range bs {
		n <<= 1
		if b == '#' {
			n += 1
		}
	}
	return n
}

func parseButton(s string) set.Set[int] {
	s = s[1 : len(s)-1]
	si := set.NewFromList(aoc.StringToInts(s))
	return si
}

func indicesToInt(ss set.Set[int]) int {
	s := ss.Copy()
	n := 0
	for !s.IsEmpty() {
		v, err := s.Take()
		if err != nil {
			panic("take from empty set failed")
		}
		mask := 1 << v
		n |= mask
	}
	return n
}

func machineFromString(s string) *machine {
	bits := strings.Split(s, " ")
	wanted := parseWanted(bits[0])
	buttonBits := bits[1 : len(bits)-1]
	buttons := fun.Map(parseButton, buttonBits)
	wiring := fun.Map(indicesToInt, buttons)
	jStr := bits[len(bits)-1]
	joltages := aoc.StringToInts(jStr[1 : len(jStr)-1])
	fmt.Printf("----\n")
	return &machine{
		wanted:    wanted,
		wiring:    wiring,
		buttons:   buttons,
		joltages:  joltages,
		numLights: len(bits[0]) - 2,
	}
}

func popcount(n int) int {
	c := 0
	for n > 0 {
		if n&1 != 0 {
			c += 1
		}
		n >>= 1
	}
	return c
}

// bools[0] is least significant bit of n
func intToBools(n int, nbits int) ([]bool, int) {
	bs := make([]bool, nbits)
	mask := 1
	popcount := 0
	for i := 0; i < nbits; i++ {
		if n&mask != 0 {
			bs[i] = true
			popcount++
		}
		mask <<= 1
	}
	return bs, popcount
}

func (m machine) getPress(bs []bool) int {
	result := 0
	for i, b := range bs {
		if b {
			result ^= m.wiring[i]
		}
	}
	return result
}

func (m machine) bestPressPart1() int {
	bestPresses := m.numButtons() + 1
	for n := range 1 << m.numButtons() {
		bs, presses := intToBools(n, m.numButtons())
		result := m.getPress(bs)
		if result == m.wanted {
			if presses < bestPresses {
				bestPresses = presses
			}
		}
	}
	if bestPresses == m.numButtons()+1 {
		panic("can't find result")
	}
	return bestPresses
}

func (m machine) bestPressPart2() int {
	buttonsForJoltage := make([]set.Set[int], m.numJoltages())
	for ji := range buttonsForJoltage {
		buttonsForJoltage[ji] = set.New[int]()
		for bi := range m.buttons {
			if m.buttons[bi].Contains(ji) {
				buttonsForJoltage[ji].Insert(bi)
			}
		}
	}
	maxForButtons := make([]int, m.numButtons())
	for bi, b := range m.buttons {
		var joltages []int
		for _, ji := range b.ToList() {
			joltages = append(joltages, m.joltages[ji])
		}
		maxForButtons[bi] = fun.Min(joltages)
	}

	fmt.Printf("buttons: %s\n", m.buttons)
	for ji := range buttonsForJoltage {
		fmt.Printf("B4J %d: %v\n", ji, buttonsForJoltage[ji])
	}
	for bi, bmax := range maxForButtons {
		fmt.Printf("%d mx %d\n", bi, bmax)
	}
	return 0
}

func (d *Day10) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")
	ms := fun.Map(machineFromString, lines)
	for _, m := range ms {
		fmt.Printf("%s\n", m)
	}
	presses := fun.Map(func(m *machine) int { return m.bestPressPart1() }, ms)
	fmt.Printf("presses: %v\n", presses)
	fmt.Printf("Part 1: %d\n", fun.Sum(presses))
	presses = fun.Map(func(m *machine) int { return m.bestPressPart2() }, ms)
	fmt.Printf("presses: %v\n", presses)
	fmt.Printf("Part 2: %d\n", fun.Sum(presses))

	return nil
}
