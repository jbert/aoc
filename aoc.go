package aoc

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"github.com/jbert/aoc/fun"
)

type Day interface {
	Run(out io.Writer, lines []string) error
}

func Run(d Day, day int, test bool, out io.Writer) error {
	lines := GetLines(day, test)
	fmt.Fprintf(out, "Lines are %v\n", lines)

	err := d.Run(out, lines)
	if err != nil {
		return fmt.Errorf("Failed running day [%d]: $s", err)
	}

	return nil
}

func GetLines(day int, test bool) []string {
	fname := dataFileName(day, test)
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalf("Can't read data file [%s]: %s", fname, err)
	}
	lines := strings.Split(string(buf), "\n")
	return fun.Filter(func(s string) bool { return s != "" }, lines)
}

func dataFileName(day int, test bool) string {
	workDir := "/home/john/dev/jbert/aoc2021"
	suffix := ""
	if test {
		suffix = "-test"
	}
	return fmt.Sprintf("%s/data/day%d%s.txt", workDir, day, suffix)
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
