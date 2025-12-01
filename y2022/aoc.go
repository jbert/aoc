package y2022

import "github.com/jbert/aoc/year"

// TODO: remove once cmds/main/y2022.go deleted
type Year struct{}

func (y *Year) WorkDir() string {
	return "/home/john/dev/jbert/aoc/data/y2022"
}

var Y = year.New(2022, dayFuncs)

var dayFuncs = map[int]year.Day{
	1:  &Day1{},
	2:  &Day2{},
	3:  &Day3{},
	4:  &Day4{},
	5:  &Day5{},
	8:  &Day8{},
	9:  &Day9{},
	10: &Day10{},
	11: &Day11{},
	12: &Day12{},
	13: &Day13{},
	14: &Day14{},
	15: &Day15{},
	16: &Day16{},
	17: &Day17{},
}
