package y2022

import (
	"io"
)

type Day12 struct{ Year }

func NewDay12() *Day12 {
	d := Day12{}
	return &d
}

func (d *Day12) Run(out io.Writer, lines []string) error {
	return nil
}
