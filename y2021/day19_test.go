package y2021

import (
	"fmt"
	"sort"
	"testing"

	"github.com/jbert/aoc/pts"
	"github.com/jbert/set"
	"github.com/stretchr/testify/assert"
)

func TestScannerRotations(t *testing.T) {
	a := assert.New(t)
	l := []string{
		"--- scanner 7 ---", "1,0,0", "0,2,0", "0,0,3",
	}
	s, l := readScanner(l)
	fmt.Printf("s is %s\n", s)
	a.Equal([]string{}, l, "consumed all lines")
	a.Equal(s.id, 7, "correct scanner id")
	a.Equal(s.beacons.Size(), 3, "three beacons")
	bss := s.BeaconSets()
	a.Equal(24, len(bss), "Have 24 beaconsets")

	strs := set.New[string]()
	for _, bs := range bss {
		bsl := bs.ToList()
		sort.Slice(bsl, func(i, j int) bool {
			return bsl[i].Less(bsl[j])
		})
		fmt.Printf("%v\n", bsl)
		strs.Insert(fmt.Sprintf("%v", bsl))
		t.Logf("BS deltas:\n%v", bs.Deltas())
	}
	a.Equal(24, strs.Size(), "24 distinct sets of pts")

	//	u, _ := readScanner([]string{
	//		"--- scanner 6 ---", "1,0,0", "0,2,0", "0,0,4",
	//	})
}

func TestBaseRotations(t *testing.T) {
	a := assert.New(t)
	x := pts.P3{1, 0, 0}
	y := pts.P3{0, 1, 0}
	z := pts.P3{0, 0, 1}
	p := pts.P3{1, 2, 3}

	a.Equal(x, rotX(x), "rot around same axis - x")
	a.Equal(y, rotY(y), "rot around same axis - y")
	a.Equal(z, rotZ(z), "rot around same axis - z")

	a.Equal(pts.P3{1, -3, 2}, rotX(p), "rotX")
	a.Equal(pts.P3{3, 2, -1}, rotY(p), "rotY")
	a.Equal(pts.P3{-2, 1, 3}, rotZ(p), "rotZ")

	a.Equal(pts.P3{1, 3, -2}, c(rotX, rotX, rotX)(p), "3rotX")
	a.Equal(pts.P3{1, -3, 2}, xRots[1](p), "rotX - from array")
	a.Equal(pts.P3{1, 3, -2}, xRots[3](p), "3rotX - from array")

	x2 := c(rotX, rotX)
	a.Equal(pts.P3{1, -2, -3}, x2(p), "two-X rot")
	y2 := c(rotY, rotY)
	a.Equal(pts.P3{-1, 2, -3}, y2(p), "two-Y rot")
	z2 := c(rotZ, rotZ)
	a.Equal(pts.P3{-1, -2, 3}, z2(p), "two-Z rot")

	a.Equal(rotY(rotX(p)), c(rotY, rotX)(p), "test Compose")
	t.Logf("xTheny(p) = %s", rotY(rotX(p)))
	p1 := rotX(p)
	p2 := rotY(p1)
	a.Equal(p2, c(rotY, rotX)(p), "test Compose test")

	a.Equal(pts.P3{2, -3, -1}, p2, "xTheny by hand")

	a.Equal(c(rotY, rotY, rotX, rotX)(p), c(rotZ, rotZ)(p), "yyxx == zz")

	xTheny := c(rotY, rotX)
	a.Equal(pts.P3{2, -3, -1}, xTheny(p), "xTheny")
}
