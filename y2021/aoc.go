package y2021

import "github.com/jbert/aoc/year"

// TODO: remove once cmds/main/y2022.go deleted
type Year struct{}

func (y *Year) WorkDir() string {
	return "/home/john/dev/jbert/aoc/data/y2021"
}

var Y = year.New(2021, dayFuncs)

var dayFuncs = map[int]year.Day{
	12: &Day12{},
	15: &Day15{},
	18: &Day18{},
	19: &Day19{},
}
