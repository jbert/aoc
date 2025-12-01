package y2021

//go:generate ../generate-dayfuncs.sh

import "github.com/jbert/aoc/year"

var Y = year.New(2021, dayFuncs)
