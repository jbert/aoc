package aoc

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/jbert/aoc/grid"
)

type Day interface {
	WorkDir() string
	Run(out io.Writer, lines []string) error
}

func Run(d Day, day int, test bool, out io.Writer) error {
	lines := GetLines(d, day, test)
	if test {
		fmt.Fprintf(out, "Lines are %v\n", lines)
	}

	err := d.Run(out, lines)
	if err != nil {
		return fmt.Errorf("Failed running day [%d]: $s", err)
	}

	return nil
}

type BaseDay struct {
}

func GetLines(d Day, day int, test bool) []string {
	fname := dataFileName(d.WorkDir(), day, test)
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalf("Can't read data file [%s]: %s", fname, err)
	}
	lines := strings.Split(string(buf), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}

func IntAbs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func MustAtoi(l string) int {
	n, err := strconv.Atoi(l)
	if err != nil {
		panic(fmt.Sprintf("Line [%s] failed to be a number: %s", l, err))
	}
	return n
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
