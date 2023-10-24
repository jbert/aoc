package y2022

import (
	"testing"

	"github.com/jbert/aoc/pts"
	"github.com/stretchr/testify/assert"
)

func TestD17Rect(t *testing.T) {
	a := assert.New(t)

	var c chamber
	r := c.rectAt(0, 0, 2, 2)
	a.Equal([][]bool{{false, false}, {false, false}}, r)
	r = c.rectAt(-1, 0, 2, 2)
	a.Equal([][]bool{{true, false}, {true, false}}, r)
	r = c.rectAt(0, -1, 2, 2)
	a.Equal([][]bool{{false, false}, {true, true}}, r)

	r = c.rectAt(chamberWidth-1, 0, 2, 2)
	a.Equal([][]bool{{false, true}, {false, true}}, r)

	q := newPiece([]string{"####"})

	q.pos = c.startPos(q)
	worked := q.tryMove(pts.P2{+1, 0}, c)
	a.Equal(true, worked)

	q.pos = pts.P2{chamberWidth - q.w, 0}
	worked = q.tryMove(pts.P2{+1, 0}, c)
	a.Equal(false, worked)

	q.pos = pts.P2{chamberWidth - q.w - 2, 0}
	worked = q.tryMove(pts.P2{+1, 0}, c)
	a.Equal(true, worked)

	p := newPiece([]string{".#.", "###", ".#."})

	p.pos = pts.P2{0, 0}
	worked = p.tryMove(pts.P2{-1, 0}, c)
	a.Equal(false, worked)

	p.pos = pts.P2{0, 0}
	worked = p.tryMove(pts.P2{0, -1}, c)
	a.Equal(false, worked)

	p.pos = pts.P2{0, 0}
	worked = p.tryMove(pts.P2{0, +1}, c)
	a.Equal(true, worked)

	p.pos = pts.P2{0, 0}
	worked = p.tryMove(pts.P2{+1, 0}, c)
	a.Equal(true, worked)

	p.pos = pts.P2{chamberWidth - p.w - 1, 0}
	worked = p.tryMove(pts.P2{+1, 0}, c)
	a.Equal(true, worked)

	p.pos = pts.P2{chamberWidth - p.w, 0}
	worked = p.tryMove(pts.P2{+1, 0}, c)
	a.Equal(false, worked)

	p.pos = pts.P2{0, 0}
	c.addPiece(p)

	p = newPiece([]string{".#.", "###", ".#."})
	p.pos = pts.P2{3, 0}
	worked = p.tryMove(pts.P2{+1, 0}, c)
	a.Equal(true, worked)

	p.pos = pts.P2{3, 0}
	worked = p.tryMove(pts.P2{-1, 0}, c)
	a.Equal(false, worked)

}
