package y2022

import (
	"fmt"
	"io"
	"strings"

	"github.com/jbert/aoc"
	"github.com/jbert/aoc/fun"
	"github.com/jbert/aoc/pts"
	"github.com/jbert/aoc/set"
)

type Day9 struct{ Year }

func NewDay9() *Day9 {
	d := Day9{}
	return &d
}

type Move struct {
	dir   pts.P2
	steps int
}

func (m Move) String() string {
	return fmt.Sprintf("%s: %d\n", m.dir, m.steps)
}

type Snake struct {
	head, tail pts.P2
}

func (s Snake) String() string {
	return fmt.Sprintf("H %s T %s", s.head, s.tail)
}

func (s *Snake) Apply(m Move) []pts.P2 {
	var tails []pts.P2
	for i := 0; i < m.steps; i++ {
		s.head = s.head.Add(m.dir)
		s.updateTail()
		tails = append(tails, s.tail)
	}
	return tails
}

func (s *Snake) updateTail() {
	d := s.head.Sub(s.tail)
	if aoc.IntAbs(d.X) >= 3 || aoc.IntAbs(d.Y) >= 3 {
		panic(fmt.Sprintf("Head too far from tail: %v", s))
	}
	switch {
	// Two steps away becomes 1, 1 step away becomes zero
	case d.X == 0 || d.Y == 0 || (aoc.IntAbs(d.X) == 1 && aoc.IntAbs(d.Y) == 1):
		d = d.Div(2)
		// Otherwise we want to round 2's to 1's
	case aoc.IntAbs(d.X) == 2:
		d.X /= 2
	case aoc.IntAbs(d.Y) == 2:
		d.Y /= 2
	default:
		panic(fmt.Sprintf("wtf: s [%s] d [%s]", s, d))
	}
	s.tail = s.tail.Add(d)
}

func (d *Day9) Run(out io.Writer, lines []string) error {
	moves := fun.Map(lineToMove, lines)
	fmt.Printf("%v\n", moves)

	zero := pts.P2{0, 0}
	snake := Snake{zero, zero}

	tailBeen := set.New[pts.P2]()
	tailBeen.Insert(zero)
	for _, move := range moves {
		fmt.Printf("Before S: %s\tM: %s\n", snake, move)
		tails := snake.Apply(move)
		tailBeen.InsertList(tails)
		fmt.Printf("After  S: %s\tM: %s\n", snake, move)
	}
	fmt.Printf("Part 1: %d\n", tailBeen.Size())
	return nil
}

func lineToMove(l string) Move {
	bits := strings.Split(l, " ")
	var dir pts.P2
	switch bits[0] {
	case "L":
		dir = pts.L
	case "R":
		dir = pts.R
	case "U":
		dir = pts.U
	case "D":
		dir = pts.D
	default:
		panic(fmt.Sprintf("Unknown dir: %s", bits[0]))
	}
	return Move{
		dir:   dir,
		steps: aoc.MustAtoi(bits[1]),
	}
}
