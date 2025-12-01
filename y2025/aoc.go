package y2025

import (
	"github.com/jbert/aoc/year"
)

var Y = year.New(2025, dayFuncs)

var dayFuncs = map[int]year.Day{
	1: &Day1{},
}
