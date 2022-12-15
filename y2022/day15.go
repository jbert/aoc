package y2022

import (
	"fmt"
	"io"
	"strings"

	"github.com/jbert/aoc/fun"
	"github.com/jbert/aoc/num"
	"github.com/jbert/aoc/pts"
)

type Day15 struct{ Year }

func NewDay15() *Day15 {
	d := Day15{}
	return &d
}

func (d *Day15) Run(out io.Writer, lines []string) error {
	sbs := fun.Map(lineToSB, lines)
	for _, sb := range sbs {
		fmt.Printf("%s\n", sb)
	}
	return nil
}

type SB struct {
	sensor pts.P2
	beacon pts.P2
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
