package y2025

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/jbert/aoc/year"
	"github.com/jbert/fun"
)

type Day2 struct{ year.Year }

type Range struct {
	lo int
	hi int
}

func (r Range) String() string {
	return fmt.Sprintf("%d->%d", r.lo, r.hi)
}

func (r Range) SimpleRanges() []Range {
	evenLo := upToEvenDigits(r.lo)
	evenHi := downToEvenDigits(r.hi)
	fmt.Printf("r %s evenLo %d evenHi %d\n", r, evenLo, evenHi)
	var ranges []Range
	if evenLo > evenHi {
		return ranges
	}

	digLo := numDigits(evenLo)
	digHi := numDigits(evenHi)
	if digLo != digHi {
		panic(fmt.Sprintf("big range: [%d, %d] %s", evenLo, evenHi, r))
	}
	ranges = append(ranges, Range{evenLo, evenHi})
	return ranges
}

func (r Range) isSimple() bool {
	digLo := numDigits(r.lo)
	digHi := numDigits(r.hi)
	return isEven(digLo) && isEven(digHi) && digLo == digHi
}

func (r Range) InvalidInSimpleRange() ([]int, error) {
	if !r.isSimple() {
		return nil, fmt.Errorf("range %s not simple", r)
	}
	lo := max(firstHalf(r.lo), secondHalf(r.lo))
	hi := min(firstHalf(r.hi), secondHalf(r.hi))
	if hi < lo {
		return nil, nil
	}
	fmt.Printf("lo %d hi %d\n", lo, hi)
	nums := fun.Iota(lo, hi-lo+1)
	numToId := func(n int) int {
		nDig := numDigits(n)
		factor := pow10(nDig + 1)
		return n * (factor + 1)
	}
	ids := fun.Map(numToId, nums)
	// return []int{hi - lo + 1}, nil
	return ids, nil
}

func firstHalf(n int) int {
	nDig := numDigits(n)
	if !isEven(nDig) {
		panic(fmt.Sprintf("firstHalf on odd-dig number %d", n))
	}
	return (n - secondHalf(n)) / pow10(1+nDig/2)
}

func secondHalf(n int) int {
	nDig := numDigits(n)
	if !isEven(nDig) {
		panic(fmt.Sprintf("secondHalf on odd-dig number %d", n))
	}
	modN := pow10(1 + nDig/2)
	sh := n % modN
	// fmt.Printf("SH: %d nDig %d pow10(nDig/2) %d sh %d\n", n, nDig, modN, sh)
	return sh
}

func upToEvenDigits(n int) int {
	nDig := numDigits(n)
	if isEven(nDig) {
		return n
	}
	return pow10(nDig + 1)
}

func downToEvenDigits(n int) int {
	nDig := numDigits(n)
	if isEven(nDig) {
		return n
	}
	return makeHi(nDig)
}

// The power of then with this number of digits
func pow10(nDig int) int {
	return int(math.Pow10(nDig - 1))
}

func makeHi(nDig int) int {
	return pow10(nDig) - 1
}

func numDigits(n int) int {
	return int(math.Log10(float64(n))) + 1
}

func isEven(n int) bool {
	return n%2 == 0
}

func parseRange(s string) (Range, error) {
	r := Range{}
	bits := strings.Split(s, "-")
	if len(bits) != 2 {
		return r, fmt.Errorf("can't parse [%s]: not two bits", s)
	}
	lo, err := strconv.Atoi(bits[0])
	if err != nil {
		return r, fmt.Errorf("can't parse [%s] as int: %w", bits[0], err)
	}
	hi, err := strconv.Atoi(bits[1])
	if err != nil {
		return r, fmt.Errorf("can't parse [%s] as int: %w", bits[0], err)
	}
	return Range{lo, hi}, nil
}

func (d *Day2) Run(out io.Writer, lines []string) error {
	line := strings.ReplaceAll(lines[0], " ", "")
	rangeStrs := strings.Split(line, ",")
	ranges, err := fun.ErrMap(parseRange, rangeStrs)
	if err != nil {
		return fmt.Errorf("can't parse: %w", err)
	}
	sRangesNested := fun.Map(func(r Range) []Range { return r.SimpleRanges() }, ranges)
	sRanges := fun.Flatten(sRangesNested)
	nestedCounts, err := fun.ErrMap(func(r Range) ([]int, error) {
		ids, err := r.InvalidInSimpleRange()
		fmt.Fprintf(out, "r [%s] count %v err %s\n", r, ids, err)
		return ids, err
	}, sRanges)
	if err != nil {
		return fmt.Errorf("can't count simple ranges: %w", err)
	}
	counts := fun.Flatten(nestedCounts)
	fmt.Fprintf(out, "Part 1: %d\n", fun.Sum(counts))

	return nil
}
