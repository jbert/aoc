package y2025

import (
	"fmt"
	"io"
)

type Day1 struct{ Year }

func NewDay1() *Day1 {
	d := Day1{}
	return &d
}

func (d *Day1) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")
	fmt.Printf("%v\n", lines)

	return nil
}
