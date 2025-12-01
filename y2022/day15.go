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

type Day15 struct{}

func NewDay15() *Day15 {
	d := Day15{}
	return &d
}

func (d *Day15) Run(out io.Writer, lines []string) error {
	sbs := fun.Map(lineToSB, lines)
	for _, sb := range sbs {
		fmt.Printf("%s\n", sb)
	}

	// Make a big enough box
	minX := fun.Min(fun.Map(minX, sbs))
	maxX := fun.Max(fun.Map(maxX, sbs))
	fmt.Printf("X: %d - %d\n", minX, maxX)

	examineY := 2000000
	if len(sbs) == 14 { // Test data
		examineY = 10
	}
	ps := fun.Map(func(x int) pts.P2 { return pts.P2{x, examineY} }, fun.Iota(minX, maxX-minX+1))
	//	fmt.Printf("PS %v\n", ps)
	inRange := func(p pts.P2) []bool {
		return fun.Map(func(sb SB) bool {
			if sb.beacon.Equals(p) {
				return false
			}
			if sb.sensor.Equals(p) {
				return false
			}
			inRange := sb.inRange(p)
			//			fmt.Printf("sb [%s] to %s\t[%d vs %d], inRange %v\n", sb, p, sb.dist(), sb.sensor.Sub(p).ManhattanLength(), inRange)
			return inRange
		}, sbs)
	}
	//	fmt.Printf("%v\n", inRange(pts.P2{0, 10}))
	fmt.Printf("Part 1: %d\n", len(fun.Filter(fun.Id[bool], fun.Map(fun.AnyBool, fun.Map(inRange, ps)))))

	fmt.Printf("\n\n")

	var sbps []SBPair
	for i, sba := range sbs {
		for _, sbb := range sbs[i+1:] {
			sd := sba.sensor.Sub(sbb.sensor).ManhattanLength()
			dd := sba.dist() + sbb.dist()
			if sd-dd == 2 {
				fmt.Printf("%s, %s:\t%d %d : %d\n", sba.sensor, sbb.sensor, sd, dd, sd-dd)
				sbp := SBPair{sba, sbb}
				sbps = append(sbps, sbp)
			}
		}
	}
	fmt.Printf("SBPairs: %v\n", sbps)

	if len(sbps) != 2 {
		panic("wtf")
	}
	possibles := sbps[0].BoundaryOverlap().Intersect(sbps[1].BoundaryOverlap())
	fmt.Printf("possibles: %v\n", possibles)
	p := possibles.ToList()[0]
	fmt.Printf("Part 2: %d\n", 4000000*p.X+p.Y)

	return nil
}

type SBPair struct {
	a, b SB
}

func (sbp SBPair) BoundaryOverlap() set.Set[pts.P2] {
	sa := set.NewFromList(sbp.a.boundary())
	sb := set.NewFromList(sbp.b.boundary())
	return sa.Intersect(sb)
}

type SB struct {
	sensor pts.P2
	beacon pts.P2
}

func (sb SB) dist() int {
	return sb.sensor.Sub(sb.beacon).ManhattanLength()
}

func (sb SB) inRange(p pts.P2) bool {
	return sb.sensor.Sub(p).ManhattanLength() <= sb.dist()
}

func (sb SB) boundary() []pts.P2 {
	var boundary []pts.P2
	d := sb.dist() + 1
	for i := 0; i <= d; i++ {
		boundary = append(boundary, sb.sensor.Add(pts.P2{i, d - i}))
		boundary = append(boundary, sb.sensor.Add(pts.P2{d - i, -i}))
		boundary = append(boundary, sb.sensor.Add(pts.P2{-i, -d + i}))
		boundary = append(boundary, sb.sensor.Add(pts.P2{-d + i, +i}))
	}
	return boundary
}

func minX(sb SB) int {
	d := sb.dist()
	xs := []int{sb.sensor.X, sb.beacon.X, sb.sensor.X - d, sb.beacon.X - d}
	return fun.Min(xs)
}

func maxX(sb SB) int {
	d := sb.dist()
	xs := []int{sb.sensor.X, sb.beacon.X, sb.sensor.X + d, sb.beacon.X + d}
	return fun.Max(xs)
}

func (sb SB) String() string {
	return fmt.Sprintf("S %s B %s", sb.sensor, sb.beacon)
}

func lineToSB(l string) SB {
	sb := SB{}
	i := strings.Index(l, "x=")
	l = l[i:]
	i = strings.Index(l, ":")
	sb.sensor = strToPt(l[:i])
	l = l[i:]
	i = strings.Index(l, "x=")
	sb.beacon = strToPt(l[i:])
	return sb
}

// x=<num>, y=<num>
func strToPt(s string) pts.P2 {
	var p pts.P2
	s = s[2:]
	t := strings.Index(s, ",")
	p.X = num.MustAtoi(s[:t])
	s = s[t+4:]
	p.Y = num.MustAtoi(s)
	return p
}
