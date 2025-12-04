package pts

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jbert/aoc/num"
)

type P2 struct {
	X, Y int
}

var (
	N  P2 = P2{0, +1}
	E     = P2{+1, 0}
	S     = P2{0, -1}
	W     = P2{-1, 0}
	NE    = N.Add(E)
	SE    = S.Add(E)
	SW    = S.Add(W)
	NW    = N.Add(W)

	L = W
	R = E
	U = N
	D = S

	NESW       = []P2{N, E, S, W}
	NEIGHBOURS = []P2{N, E, S, W, NE, SE, SW, NW}
)

func (p P2) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func P2FromString(s string) P2 {
	bits := strings.Split(s, ",")
	if len(bits) != 2 {
		panic(fmt.Sprintf("Don't have 2 bits: [%s]", s))
	}
	var p P2
	var err error

	p.X, err = strconv.Atoi(bits[0])
	if err != nil {
		panic(fmt.Sprintf("Bad X coord [%s]: %s", bits[0], err))
	}
	p.Y, err = strconv.Atoi(bits[1])
	if err != nil {
		panic(fmt.Sprintf("Bad Y coord [%s]: %s", bits[1], err))
	}
	return p
}

func (p P2) ManhattanLength() int {
	return num.IntAbs(p.X) + num.IntAbs(p.Y)
}

func (p P2) IsZero() bool {
	return p.X == 0 && p.Y == 0
}

func (p P2) Add(q P2) P2 {
	return P2{
		X: p.X + q.X,
		Y: p.Y + q.Y,
	}
}

func (p P2) Sub(q P2) P2 {
	return P2{
		X: p.X - q.X,
		Y: p.Y - q.Y,
	}
}

func (p P2) Less(q P2) bool {
	if p.X < q.X {
		return true
	}
	if p.Y < q.Y {
		return true
	}
	return false
}

func (p P2) Div(n int) P2 {
	return P2{p.X / 2, p.Y / 2}
}

func (p P2) Equals(q P2) bool {
	return p.X == q.X && p.Y == q.Y
}
