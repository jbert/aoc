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
	lgs := aoc.LineGroups(strings.Split(pieceStr, "\n"))
	pieces := fun.Map(newPiece, aoc.LineGroups(strings.Split(pieceStr, "\n")))
	//	for _, p := range pieces {
	//		fmt.Printf("%s\n", p)
	//	}

	jets := lines[0]
	fmt.Printf("JETS: %s\n", jets)
	iJet := 0
	downDir := pts.P2{0, -1}

	var c chamber

	numStopped := 0
	iPiece := 0
	for numStopped <= 2022 {
		p := newPiece(lgs[iPiece])
		iPiece++
		iPiece = iPiece % len(pieces)

		//		fmt.Printf("%s\n", p)
		p.pos = c.startPos(p)
		p.moving = true
		c.addPiece(p)
		//		fmt.Printf("%s\n", c.String())
		for p.moving {
			jetChar := jets[iJet]
			iJet++
			iJet = iJet % len(jets)

			jetDir := pts.P2{+1, 0}
			if jetChar == '<' {
				jetDir = pts.P2{-1, 0}
			}
			//			fmt.Printf("%s", c.String())
			worked := p.tryMove(jetDir, c)
			//			fmt.Printf("Push [%s] worked [%v]\n\n", jetDir, worked)
			//			fmt.Printf("%s", c.String())
			worked = p.tryMove(downDir, c)
			//			fmt.Printf("Drop worked [%v]\n\n", worked)
			if !worked {
				p.moving = false
				numStopped++
			}
		}
	}
	fmt.Printf("Stopped height: %d\n", c.highestTop(false)-1)

	return nil
}

type chamber []*piece

func (c chamber) startPos(p *piece) pts.P2 {
	/*
		Each rock appears so that its left edge is two units away from the left
		wall and its bottom edge is three units above the highest rock in the room
		(or the floor, if there isn't one)
	*/
	start := pts.P2{X: 2, Y: c.highestTop(false) + 3 + 1}
	return start
}

func (c *chamber) addPiece(p *piece) {
	//	fmt.Printf("Adding:\n%s\n", p)
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
	for j := c.highestTop(true); j >= 0; j-- {
		fmt.Fprintf(b, "|")
		for i := 0; i < chamberWidth; i++ {
			c := c.charAt(i, j)
			fmt.Fprintf(b, "%c", c)
		}
		fmt.Fprintf(b, "|\n")
	}
	fmt.Fprintf(b, "+")
	for i := 0; i < chamberWidth; i++ {
		fmt.Fprintf(b, "-")
	}
	fmt.Fprintf(b, "+\n")
	fmt.Fprintf(b, "Highest stopped: %d\n", c.highestTop(false))
	return b.String()
}

func (c chamber) highestTop(includeMoving bool) int {
	if len(c) == 0 {
		return -1
	}
	h := 0
	for _, p := range c {
		if includeMoving || !p.moving {
			h = max(h, p.top())
		}
	}
	return h
}

func (c chamber) rectAt(x, y, w, h int) [][]bool {
	rect := make([][]bool, h)
	for j := 0; j < h; j++ {
		rect[j] = make([]bool, w)
		for i := 0; i < w; i++ {
			if i+x > chamberWidth-1 || i+x < 0 || j+y < 0 {
				rect[j][i] = true
				continue
			}
			char := c.charAt(i+x, j+y)
			rect[j][i] = char == '#'
		}
	}
	return rect
}

func (c chamber) charAt(x, y int) byte {
	for _, p := range c {
		c := p.charAt(x, y)
		if c != ' ' {
			return c
		}
	}
	return '.'
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
	for j, l := range lines {
		p.bits[p.h-1-j] = make([]bool, p.w)
		for i, r := range l {
			p.bits[p.h-1-j][i] = r == '#'
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
	if p.bits[y][x] {
		if p.moving {
			return '@'
		} else {
			return '#'
		}
	}
	return ' '
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

func rectIntersect(a, b [][]bool) bool {
	if len(a) != len(b) || len(a) == 0 || len(a[0]) != len(b[0]) {
		panic("wtf")
	}
	//	fmt.Printf("JB RI:\n%v\n%v\n", a, b)
	for j := range a {
		for i := range a[j] {
			if a[j][i] && b[j][i] {
				return true
			}
		}
	}
	return false
}

func (p *piece) tryMove(dir pts.P2, c chamber) bool {
	moveTo := p.pos.Add(dir)
	r := c.rectAt(moveTo.X, moveTo.Y, p.w, p.h)
	//	fmt.Printf("JB: c.rect:\n%v\n%v\n", r, p.rect())
	worked := !rectIntersect(r, p.rect())
	if worked {
		p.pos = moveTo
	}
	return worked
}

func (p *piece) rect() [][]bool {
	rect := make([][]bool, p.h)
	for j := 0; j < p.h; j++ {
		rect[j] = make([]bool, p.w)
		for i := 0; i < p.w; i++ {
			rect[j][i] = p.bits[j][i]
		}
	}
	return rect
}

func (p *piece) String() string {
	b := &strings.Builder{}
	fmt.Fprintf(b, "X %d Y %d\n", p.pos.X, p.pos.Y)
	fmt.Fprintf(b, "W %d H %d\n", p.w, p.h)
	r := '#'
	if p.moving {
		r = '@'
	}
	for j := range p.bits {
		for i := range p.bits[0] {
			if p.bits[p.h-1-j][i] {
				fmt.Fprintf(b, "%c", r)
			} else {
				fmt.Fprintf(b, " ")
			}
		}
		fmt.Fprintf(b, "\n")
	}
	return b.String()
}
