package y2022

import (
	"fmt"
	"io"
	"strings"

	"github.com/jbert/aoc"
)

type Day17 struct{ Year }

func NewDay17() *Day17 {
	d := Day17{}
	return &d
}

func (d *Day17) Run(out io.Writer, lines []string) error {
	pieces := aoc.LineGroups(strings.Split(pieceStr, "\n"))
	for _, p := range pieces {
		fmt.Printf("%s\n", p)
	}
	return nil
}

const pieceStr = `####

.#.
###
.#.

..#
..#
###

#
#
#
#

##
##`
