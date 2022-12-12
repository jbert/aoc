package pts

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jbert/aoc/num"
)

type P3 struct {
	X, Y, Z int
}

func (p P3) String() string {
	return fmt.Sprintf("%d,%d,%d", p.X, p.Y, p.Z)
}

func P3FromString(s string) P3 {
	bits := strings.Split(s, ",")
	if len(bits) != 3 {
		panic(fmt.Sprintf("Don't have 3 bits: [%s]", s))
	}
	var p P3
	var err error

	p.X, err = strconv.Atoi(bits[0])
	if err != nil {
		panic(fmt.Sprintf("Bad X coord [%s]: %s", bits[0], err))
	}
	p.Y, err = strconv.Atoi(bits[1])
	if err != nil {
		panic(fmt.Sprintf("Bad Y coord [%s]: %s", bits[1], err))
	}
	p.Z, err = strconv.Atoi(bits[2])
	if err != nil {
		panic(fmt.Sprintf("Bad Z coord [%s]: %s", bits[2], err))
	}
	return p
}

func (p P3) ManhattanLength() int {
	return num.IntAbs(p.X) + num.IntAbs(p.Y) + num.IntAbs(p.Z)
}

func (p P3) IsZero() bool {
	return p.X == 0 && p.Y == 0 && p.Z == 0
}

func (p P3) Add(q P3) P3 {
	return P3{
		X: p.X + q.X,
		Y: p.Y + q.Y,
		Z: p.Z + q.Z,
	}
}

func (p P3) Sub(q P3) P3 {
	return P3{
		X: p.X - q.X,
		Y: p.Y - q.Y,
		Z: p.Z - q.Z,
	}
}

func (p P3) Less(q P3) bool {
	if p.X < q.X {
		return true
	}
	if p.Y < q.Y {
		return true
	}
	if p.Z < q.Z {
		return true
	}
	return false
}
