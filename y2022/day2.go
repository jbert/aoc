package y2022

import (
	"fmt"
	"io"

	"github.com/jbert/fun"
)

type Day2 struct{ Year }

func NewDay2() *Day2 {
	d := Day2{}
	return &d
}

type RPS int

func (rps RPS) Beats() RPS {
	switch rps {
	case R:
		return S
	case P:
		return R
	case S:
		return P
	default:
		panic("bad rps")
	}
}

func (rps RPS) BeatenBy() RPS {
	switch rps {
	case R:
		return P
	case P:
		return S
	case S:
		return R
	default:
		panic("bad rps")
	}
}

func (rps RPS) WinsOver(other RPS) bool {
	return (rps == R && other == S) || (rps == P && other == R) || (rps == S && other == P)
}

func (rps RPS) Value() int {
	switch rps {
	case R:
		return 1
	case P:
		return 2
	case S:
		return 3
	default:
		panic("bad rps")
	}
}

func (rps RPS) String() string {
	switch rps {
	case R:
		return "Rock"
	case P:
		return "Paper"
	case S:
		return "Sciss"
	default:
		panic("bad rps")
	}
}

const (
	R RPS = iota
	P
	S
)

type Round struct {
	o, m RPS
}

func (r Round) String() string {
	return fmt.Sprintf("O %s\tM %s", r.o, r.m)
}

func (r Round) Score() int {
	wdl := func() int {
		if r.o == r.m {
			return 3
		}
		if r.m.WinsOver(r.o) {
			return 6
		}
		if r.o.WinsOver(r.m) {
			return 0
		}
		panic(fmt.Sprintf("Bad round: %s", r))
	}
	return wdl() + r.m.Value()
}

func charToRPS(c byte) RPS {
	switch c {
	case 'A':
		return R
	case 'B':
		return P
	case 'C':
		return S
	case 'X':
		return R
	case 'Y':
		return P
	case 'Z':
		return S
	default:
		panic(fmt.Sprintf("Bad char: %v", c))
	}
}

func lineToRound(l string) Round {
	return Round{
		o: charToRPS(l[0]),
		m: charToRPS(l[2]),
	}
}

type Outcome int

const (
	Lose Outcome = iota
	Draw
	Win
)

func charToOutcome(c byte) Outcome {

	switch c {
	case 'X':
		return Lose
	case 'Y':
		return Draw
	case 'Z':
		return Win
	default:
		panic(fmt.Sprintf("Bad char: %v", c))
	}
}

type RoundP2 struct {
	o       RPS
	outcome Outcome
}

func lineToRoundP2(l string) RoundP2 {
	return RoundP2{
		o:       charToRPS(l[0]),
		outcome: charToOutcome(l[2]),
	}
}

func rp2Round(rp2 RoundP2) Round {
	r := Round{o: rp2.o}
	switch rp2.outcome {
	case Draw:
		r.m = r.o
	case Win:
		r.m = r.o.BeatenBy()
	case Lose:
		r.m = r.o.Beats()
	default:
		panic("bad outcome")
	}
	return r
}

func (d *Day2) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")

	rounds := fun.Map(lineToRound, lines)

	for _, r := range rounds {
		fmt.Printf("%s: %d\n", r, r.Score())
	}
	fmt.Printf("Part 1: %d\n", fun.Sum(fun.Map(func(r Round) int { return r.Score() }, rounds)))

	rp2s := fun.Map(lineToRoundP2, lines)
	p2rounds := fun.Map(rp2Round, rp2s)
	fmt.Printf("Part 1: %d\n", fun.Sum(fun.Map(func(r Round) int { return r.Score() }, p2rounds)))

	return nil
}
