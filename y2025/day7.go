package y2025

import (
	"fmt"
	"io"
	"strings"

	"github.com/jbert/aoc/year"
	"github.com/jbert/fun"
	"github.com/jbert/set"
)

type Day7 struct{ year.Year }

func (d *Day7) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")

	sPos := strings.Index(lines[0], "S")
	beams := []int{sPos}
	splits := 0

	bts := make([]int, len(lines[0]))
	bts[sPos] = 1

	for j := range lines {
		if j == 0 {
			continue
		}
		newBeams := set.New[int]()
		nbts := make([]int, len(lines[0]))
		for _, beam := range beams {
			if lines[j][beam] == '^' {
				newBeams.Insert(beam - 1)
				newBeams.Insert(beam + 1)
				nbts[beam-1] += bts[beam]
				nbts[beam+1] += bts[beam]
				splits += 1
			} else {
				newBeams.Insert(beam)
				nbts[beam] += bts[beam]
			}
		}
		beams = newBeams.ToList()
		bts = nbts
	}
	fmt.Printf("Part 1: %d\n", splits)

	fmt.Printf("Part 2: %d\n", fun.Sum(bts))

	return nil
}
