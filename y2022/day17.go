package y2022

import (
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/jbert/aoc"
	"github.com/jbert/aoc/pts"
	"github.com/jbert/fun"
)

type Day17 struct{ Year }

func NewDay17() *Day17 {
	d := Day17{}
	return &d
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

const chamberWidth = 7

// All coords are (0,0) at bottom left

func (d *Day17) Run(out io.Writer, lines []string) error {
	pieces := fun.Map(newPiece, aoc.LineGroups(strings.Split(pieceStr, "\n")))
	//	for _, p := range pieces {
	//		fmt.Printf("%s\n", p)
	//	}

	jets := lines[0]
	//	fmt.Printf("JETS: %s\n", jets)
	iJet := 0

	var c chamber

	for i := 1; i < 10; i++ {
		p := pieces[(i-1)%len(pieces)].Copy()
		fmt.Printf("%s\n", p)
		p.pos = c.startPos(p)
		p.moving = true
		c.addPiece(p)
		for p.moving {
			dir := jets[iJet]
			iJet++
			iJet = iJet % (len(jets) - 1)

			p.applyJet(dir, chamberWidth)
			p.applyGravity(c)

			fmt.Printf("%s\n", c.String())
		}
	}

	return nil
}

type chamber []*piece

func (c chamber) startPos(p *piece) pts.P2 {
	/*
		Each rock appears so that its left edge is two units away from the left
		wall and its bottom edge is three units above the highest rock in the room
		(or the floor, if there isn't one)
	*/
	start := pts.P2{X: 2, Y: c.highestStopped() + 3}
	return start
}

func (c *chamber) addPiece(p *piece) {
	fmt.Printf("Adding:\n%s\n", p)
	*c = append(*c, p)
	c.sort()
}

func (c *chamber) sort() {
	ps := []*piece(*c)
	slices.SortFunc(ps, func(a, b *piece) int {
		if a.bottom() < b.bottom() {
			return -1
		}
		if a.bottom() > b.bottom() {
			return +1
		}
		if a.left() < b.left() {
			return -1
		}
		if a.left() > b.left() {
			return +1
		}
		return 0
	})
	cc := chamber(ps)
	c = &cc
}

func (c chamber) String() string {
	b := &strings.Builder{}
	for j := c.highestStopped() + 6; j >= 0; j-- {
		fmt.Fprintf(b, "|")
		for i := 0; i < chamberWidth; i++ {
			printed := false
			for _, p := range c {
				c := p.charAt(i, j)
				if c != ' ' {
					fmt.Fprintf(b, "%c", p.charAt(i, j))
					printed = true
				}
			}
			if !printed {
				fmt.Fprintf(b, "%c", '.')
			}
		}
		fmt.Fprintf(b, "|\n")
	}
	fmt.Fprintf(b, "+")
	for i := 0; i < chamberWidth; i++ {
		fmt.Fprintf(b, "-")
	}
	fmt.Fprintf(b, "+\n")
	fmt.Fprintf(b, "Highest stopped: %d\n", c.highestStopped())
	return b.String()
}

func (c chamber) highestStopped() int {
	if len(c) == 0 {
		return 0
	}
	h := 0
	for _, p := range c {
		if !p.moving {
			h = max(h, p.top())
		}
	}
	return h
}

type piece struct {
	pos    pts.P2
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

func (p *piece) charAt(x, y int) byte {
	x = x - p.pos.X
	y = y - p.pos.Y

	if y < 0 || y >= p.h {
		return ' '
	}
	if x < 0 || x >= p.w {
		return ' '
	}
	if p.bits[(p.h-1)-y][x] {
		if p.moving {
			return '@'
		} else {
			return '#'
		}
	}
	return ' '
}

func (p *piece) Copy() *piece {
	// Shallow copy is fine, the 'bits' array is immutable
	q := *p
	return &q
}

func (p *piece) left() int {
	return p.pos.X
}

func (p *piece) right() int {
	return p.pos.X + p.w - 1
}

func (p *piece) top() int {
	return p.pos.Y + p.h - 1
}

func (p *piece) bottom() int {
	return p.pos.Y
}

func (p *piece) applyGravity(c chamber) {
	fmt.Printf("PB: %d HS %d\n", p.bottom(), c.highestStopped())
	if p.bottom() <= c.highestStopped() {
		fmt.Printf("Stopping\n")
		p.moving = false
		return
	}
	p.pos.Y--
}

func (p *piece) applyJet(dir byte, chamberWidth int) {
	switch dir {
	case '>':
		if p.right() == chamberWidth-1 {
			return
		}
		p.pos.X++
		return
	case '<':
		if p.left() == 0 {
			return
		}
		p.pos.X--
		return
	default:
		panic("wtf")
	}
}

func (p *piece) String() string {
	b := &strings.Builder{}
	fmt.Fprintf(b, "X %d Y %d\n", p.pos.X, p.pos.Y)
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
