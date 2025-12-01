package y2025

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/jbert/fun"
)

type Day1 struct{ Year }

func NewDay1() *Day1 {
	d := Day1{}
	return &d
}

type Pos int

func (p Pos) Perform(t Turn) Pos {
	v := Pos(int(p) + int(t))
	return v % 100
}

type Turn int

func (t Turn) String() string {
	it := int(t)
	prefix := "R"
	if it < 0 {
		prefix = "L"
		it *= -1
	}
	return fmt.Sprintf("%s%d", prefix, it)
}

func lineToTurn(l string) (Turn, error) {
	if l == "" {
		return 0, errors.New("Empty turn string")
	}
	dir := l[0]
	var sign int
	if dir == 'L' {
		sign = -1
	} else if dir == 'R' {
		sign = 1
	} else {
		return 0, fmt.Errorf("unknown direction [%c]", dir)
	}
	step, err := strconv.Atoi(l[1:])
	if err != nil {
		return 0, fmt.Errorf("can't parse as int [%s]: %w", l[1:], err)
	}
	if step == 0 {
		return 0, fmt.Errorf("zero step - check invariants")
	}
	return Turn(sign * step), nil
}

func (d *Day1) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")

	pos := Pos(50)
	turns, err := fun.ErrMap(lineToTurn, lines)
	if err != nil {
		return fmt.Errorf("can't map lines to turns: %w", err)
	}
	// fmt.Fprintf(out, "pos %d: %v\n", pos, turns)
	numZeros := 0
	fmt.Fprintf(out, "The dial starts by pointing at %d.\n", pos)
	for _, turn := range turns {
		pos = pos.Perform(turn)
		fmt.Fprintf(out, "The dial is rotated %s by to point at %d.\n", turn, pos)
		if pos == 0 {
			numZeros++
		}
	}
	fmt.Fprintf(out, "password %d\n", numZeros)

	return nil
}
