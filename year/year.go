package year

import (
	"fmt"
	"io"
)

type Year struct {
	numYear int
	days    map[int]Day
}

type Day interface {
	Run(io.Writer, []string) error
}

type DayFunc func(io.Writer, []string) error

func New(numYear int, days map[int]Day) Year {
	return Year{numYear, days}
}

func (y Year) WorkDir() string {
	return fmt.Sprintf("/home/john/dev/jbert/aoc/data/y%d", y.numYear)
}

func Missing(dayNum int) DayFunc {
	return func(out io.Writer, lines []string) error {
		return fmt.Errorf("no function for day %d", dayNum)
	}
}

func (y Year) DayFunc(numDay int) DayFunc {
	day, ok := y.days[numDay]
	if !ok {
		return Missing(numDay)
	}
	return func(out io.Writer, lines []string) error {
		return day.Run(out, lines)
	}
}
