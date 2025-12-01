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
type Turn int

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
	return Turn(sign * step), nil
}

func (d *Day1) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")

	pos := Pos(50)
	turns, err := fun.ErrMap(lineToTurn, lines)
	if err != nil {
		return fmt.Errorf("can't map lines to turns: %w", err)
	}
	fmt.Fprintf(out, "pos %d: %v\n", pos, turns)

	return nil
}
