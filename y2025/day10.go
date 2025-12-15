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
	buttons   []button
	joltages  []int
	numLights int
}

func (m machine) String() string {
	return fmt.Sprintf("%0*b: %v {%v} [%d]", m.numLights, m.wanted, m.wiring, m.joltages, m.numLights)
}

func (m machine) numJoltages() int {
	// They are equal
	return m.numLights
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

func parseWiring(s string) set.Set[int] {
	s = s[1 : len(s)-1]
	return set.NewFromList(aoc.StringToInts(s))
}

func indicesToInt(indices set.Set[int]) int {
	n := 0
	for !indices.IsEmpty() {
		v, err := indices.Take()
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
	var buttons []button
	butbits := bits[1 : len(bits)-1]
	for i, l := range butbits {
		s := parseWiring(l)
		buttons = append(buttons, button{id: i, joltages: s})

	}
	jStr := bits[len(bits)-1]
	joltages := aoc.StringToInts(jStr[1 : len(jStr)-1])
	return &machine{
		wanted:    wanted,
		wiring:    fun.Map(func(b button) int { return indicesToInt(b.joltages) }, buttons),
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

// The joltages which a button presses
type button struct {
	id       int
	joltages set.Set[int]
}

func (b button) addsTo(ij int) bool {
	return b.joltages.Contains(ij)
}

func (m machine) buttonsForJoltage(ij int) []button {
	var bs []button
	for _, b := range m.buttons {
		if b.addsTo(ij) {
			bs = append(bs, b)
		}
	}
	return bs
}

type constraint struct {
	total   int
	buttons set.Set[int]
}

type press []int

func (c constraint) compatible(ps presses) bool {
	return false
}

func (m machine) bestJoltagePress() int {
	buttonLimits := make([]int, m.numButtons)
	for ib := range buttonLimits {
		indices := intToIndices(m.wiring[ib])
		var joltages []int
	}
	return 0
}

// given a 'presses', we can:
// - return the single-elt list of this press if it matches joltages
// - return empty list and do nothing if it is joltage
// - else it is below-joltage and we can try adding each press:
//   - and recurse, getting the list of successful presses
//   - accumulate this list and return
//
// - optionally filter for least presses prior

func (m machine) getSuccessfulPress []press
func (m machine) bestPress() int {
	// we can come up with a list of max values (constraints) for buttons
	// (start with the min joltage for each button)
	//
	bestPresses := m.numButtons() + 1
	for n := range 1 << m.numButtons() {
		bs, presses := intToBools(n, m.numButtons())
		result := m.getPress(bs)
		// fmt.Printf("N %d bs [%v] result %b wanted %b\n", n, bs, result, m.wanted)
		if result == m.wanted {
			if presses < bestPresses {
				bestPresses = presses
			}
		}
	}
	if bestPresses == m.numButtons()+1 {
		panic("can't find result")
	}
	// fmt.Printf("%d ------------------\n", bestPresses)
	return bestPresses
}

func (d *Day10) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")
	ms := fun.Map(machineFromString, lines)
	for _, m := range ms {
		fmt.Printf("%s\n", m)
	}
	presses := fun.Map(func(m *machine) int { return m.bestPress() }, ms)
	fmt.Printf("presses: %v\n", presses)

	fmt.Printf("Part 1: %d\n", fun.Sum(presses))

	presses = fun.Map(func(m *machine) int { return m.bestJoltagePress() }, ms)
	fmt.Printf("presses: %v\n", presses)

	fmt.Printf("Part 2: %d\n", fun.Sum(presses))

	return nil
}
