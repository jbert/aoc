package aoc

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/jbert/aoc/grid"
	"github.com/jbert/aoc/year"
)

type Day interface {
	Run(out io.Writer, lines []string) error
}

func Run(y year.Year, dayNum int, test bool, out io.Writer) error {
	lines := getLines(y.WorkDir(), dayNum, test)
	if test {
		fmt.Fprintf(out, "lines are %v\n", lines)
	}

	f := y.DayFunc(dayNum)
	err := f(out, lines)
	if err != nil {
		return fmt.Errorf("failed running day [%d]: %s", dayNum, err)
	}

	return nil
}

func getLines(workDir string, day int, test bool) []string {
	fname := dataFileName(workDir, day, test)
	buf, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalf("Can't read data file [%s]: %s", fname, err)
	}
	lines := strings.Split(string(buf), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}

// Get lines as a char/byte grid
func ByteGrid(lines []string) grid.Grid[byte] {
	w := len(lines[0])
	h := len(lines)
	g := grid.New[byte](w, h)
	g.ForEach(func(i, j int) {
		c := lines[j][i]
		g.Set(i, j, c)
	})
	return g
}

// Get linesa as a digit grid
func IntGrid(lines []string) grid.Grid[int] {
	w := len(lines[0])
	h := len(lines)
	g := grid.New[int](w, h)
	g.ForEach(func(i, j int) {
		c := lines[j][i]
		g.Set(i, j, int(c-'0'))
	})
	return g
}

// Break lines into groups (separated by blank lines)
func LineGroups(lines []string) [][]string {
	var lgs [][]string
	var lg []string
	for _, l := range lines {
		if l == "" {
			lgs = append(lgs, lg)
			lg = make([]string, 0)
		} else {
			lg = append(lg, l)
		}
	}
	lgs = append(lgs, lg)
	lg = make([]string, 0)
	return lgs
}

func dataFileName(workDir string, day int, test bool) string {
	suffix := ""
	if test {
		suffix = "-test"
	}
	return fmt.Sprintf("%s/day%d%s.txt", workDir, day, suffix)
}

/*
 aoc-get-stream
 aoc-get-nums
 aoc-set-day
 aoc-set-test

 hash-key-add

 count-inc
 count-inc-foldl
 count-add

 list-nth
 list-partitionf

 half-cartesian-product
*/
