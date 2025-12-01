package y2022

//go:generate ../generate-dayfuncs.sh

import "github.com/jbert/aoc/year"

var Y = year.New(2022, dayFuncs)
