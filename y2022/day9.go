package y2022

import (
	"fmt"
	"io"
	"strings"

	"github.com/jbert/aoc/num"
	"github.com/jbert/aoc/pts"
	"github.com/jbert/fun"
	"github.com/jbert/set"
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
	return fmt.Sprintf("%s: %d", m.dir, m.steps)
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

func findTailUpdate(head, tail pts.P2) pts.P2 {
	d := head.Sub(tail)
	if num.IntAbs(d.X) >= 3 || num.IntAbs(d.Y) >= 3 {
		panic(fmt.Sprintf("Head too far from tail: %v -> %v", head, tail))
	}
	switch {
	// Two steps away becomes 1, 1 step away becomes zero
	case d.X == 0 ||
		d.Y == 0 ||
		(num.IntAbs(d.X) == 1 && num.IntAbs(d.Y) == 1) ||
		(num.IntAbs(d.X) == 2 && num.IntAbs(d.Y) == 2):
		d = d.Div(2)
		// Otherwise we want to round 2's to 1's
	case num.IntAbs(d.X) == 2:
		d.X /= 2
	case num.IntAbs(d.Y) == 2:
		d.Y /= 2
	default:
		panic(fmt.Sprintf("wtf: d [%s]", d))
	}
	return d
}

func (s *Snake) updateTail() {
	d := findTailUpdate(s.head, s.tail)
	s.tail = s.tail.Add(d)
}

type Rope []pts.P2

func NewRope(start pts.P2, length int) Rope {
	r := make([]pts.P2, length)
	for i := range r {
		r[i] = start
	}
	return r
}

func (r Rope) Apply(m Move) []pts.P2 {
	//	fmt.Printf("== %s ==\n", m)
	var tails []pts.P2
	for i := 0; i < m.steps; i++ {
		r[0] = r[0].Add(m.dir)
		for j := range r[1:] {
			//			fmt.Printf("RA [%d/%d] S: %s\tM: %s\n", i, m.steps, r, m)
			d := findTailUpdate(r[j], r[j+1])
			r[j+1] = r[j+1].Add(d)
		}
		//		fmt.Printf("%s\n\n", r.Display(6, 5))
		tails = append(tails, r[len(r)-1])
	}
	return tails
}

func (r Rope) Display(w, h int) string {
	var lines []string
	for j := 0; j < h; j++ {
		l := make([]byte, w)
		for i := 0; i < w; i++ {
			l[i] = '.'
			for ri := range r {
				if l[i] == '.' && r[ri].Equals(pts.P2{i, j}) {
					l[i] = byte(ri) + '0'
					if ri == 0 {
						l[i] = 'H'
					}
				}
			}
		}
		lines = append(lines, string(l))
	}
	return strings.Join(fun.Reverse(lines), "\n")
}

func (d *Day9) Run(out io.Writer, lines []string) error {
	moves := fun.Map(lineToMove, lines)
	fmt.Printf("%v\n", moves)

	zero := pts.P2{0, 0}
	snake := Snake{zero, zero}

	tailBeen := set.New[pts.P2]()
	tailBeen.Insert(zero)
	for _, move := range moves {
		//		fmt.Printf("Before S: %s\tM: %s\n", snake, move)
		tails := snake.Apply(move)
		tailBeen.InsertList(tails)
		//		fmt.Printf("After  S: %s\tM: %s\n", snake, move)
	}
	fmt.Printf("Part 1: %d\n", tailBeen.Size())

	rope := NewRope(zero, 10)
	tailBeen = set.New[pts.P2]()
	for _, move := range moves {
		//		fmt.Printf("Before S: %s\tM: %s\n", rope, move)
		tails := rope.Apply(move)
		tailBeen.InsertList(tails)
		//		fmt.Printf("%s\n\n", rope.Display(6, 6))
		//		fmt.Printf("After  S: %s\tM: %s\n", rope, move)
	}
	fmt.Printf("Part 2: %d\n", tailBeen.Size())
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
		steps: num.MustAtoi(bits[1]),
	}
}
