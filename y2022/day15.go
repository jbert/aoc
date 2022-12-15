package y2022

import (
	"io"
)

type Day15 struct{ Year }

func NewDay15() *Day15 {
	d := Day15{}
	return &d
}

func (d *Day15) Run(out io.Writer, lines []string) error {
	return nil
}
