package y2022

import (
	"fmt"
	"io"
	"strings"

	"github.com/jbert/aoc"
	"github.com/jbert/fun"
)

type Day17 struct{ Year }

func NewDay17() *Day17 {
	d := Day17{}
	return &d
}

func (d *Day17) Run(out io.Writer, lines []string) error {
	pieces := fun.Map(newPiece, aoc.LineGroups(strings.Split(pieceStr, "\n")))
	for _, p := range pieces {
		fmt.Printf("%s\n", p)
	}

	return nil
}

type piece struct {
	x, y   int
	bits   [][]bool
	w, h   int
	moving bool
}

func newPiece(lines []string) *piece {
	p := piece{
		bits: make([][]bool, len(lines)),
		h:    len(lines),
		w:    len(lines[0]),
	}
	for i, l := range lines {
		p.bits[i] = make([]bool, p.w)
		for j, r := range l {
			p.bits[i][j] = r == '#'
		}
	}
	return &p
}

func (p *piece) String() string {
	b := &strings.Builder{}
	fmt.Fprintf(b, "W %d H %d\n", p.w, p.h)
	r := '#'
	if p.moving {
		r = '@'
	}
	for i := range p.bits {
		for j := range p.bits[0] {
			if p.bits[i][j] {
				fmt.Fprintf(b, "%c", r)
			} else {
				fmt.Fprintf(b, " ")
			}
		}
		fmt.Fprintf(b, "\n")
	}
	return b.String()
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
