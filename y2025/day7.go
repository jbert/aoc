package y2025

import (
	"fmt"
	"io"
	"strings"

	"github.com/jbert/aoc/year"
	"github.com/jbert/set"
)

type Day7 struct{ year.Year }

func (d *Day7) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")

	sPos := strings.Index(lines[0], "S")
	beams := []int{sPos}
	splits := 0
	for j := range lines {
		if j == 0 {
			continue
		}
		newBeams := set.New[int]()
		for _, beam := range beams {
			if lines[j][beam] == '^' {
				newBeams.Insert(beam - 1)
				newBeams.Insert(beam + 1)
				splits += 1
			} else {
				newBeams.Insert(beam)
			}
		}
		beams = newBeams.ToList()
	}
	fmt.Printf("Part 1: %d\n", splits)
	// fmt.Printf("Part 2: %d\n", fun.Sum(joltages))

	return nil
}
