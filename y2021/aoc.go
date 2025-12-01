package y2021

import "github.com/jbert/aoc/year"

var Y = year.New(2021, dayFuncs)

var dayFuncs = map[int]year.Day{
	12: &Day12{},
	15: &Day15{},
	18: &Day18{},
	19: &Day19{},
}
