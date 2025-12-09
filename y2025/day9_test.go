package y2025

import (
	"fmt"
	"testing"

	"github.com/jbert/aoc/pts"
	"github.com/stretchr/testify/assert"
)

func baseSegs() ([]HSeg, []VSeg) {
	hsegs := []HSeg{
		// NewHSeg(pts.P2{0, 0}, pts.P2{10, 0}),
		NewHSeg(pts.P2{0, 0}, pts.P2{5, 0}),
		// NewHSeg(pts.P2{0, 0}, pts.P2{10, 0}),
		NewHSeg(pts.P2{3, 7}, pts.P2{6, 7}),
		NewHSeg(pts.P2{6, 0}, pts.P2{10, 0}),
	}
	vsegs := []VSeg{
		NewVSeg(pts.P2{0, 5}, pts.P2{0, 10}),
		NewVSeg(pts.P2{3, 0}, pts.P2{3, 8}),
		NewVSeg(pts.P2{6, 0}, pts.P2{6, 8}),
	}
	return mergeHSegs(hsegs), mergeVSegs(vsegs)
}

func TestAcceptableRect(t *testing.T) {
	a := assert.New(t)
	hs, vs := baseSegs()
	tcs := []struct {
		b        int
		l        int
		t        int
		r        int
		expected bool
	}{
		{0, 1, 0, 1, false},
		{3, 0, 6, 7, true},
	}
	for _, tc := range tcs {
		bl := pts.P2{X: tc.l, Y: tc.b}
		tr := pts.P2{X: tc.r, Y: tc.t}
		r := pts.NewRect(bl, tr)
		got := acceptableRect(r, hs, vs)
		a.Equal(tc.expected, got, fmt.Sprintf("%v", tc))
	}
}
func TestAcceptableHSegs(t *testing.T) {
	a := assert.New(t)
	hs, _ := baseSegs()
	tcs := []struct {
		x        int
		y        int
		w        int
		expected bool
	}{
		{2, 3, 5, false},
		{8, 3, 5, false},
		{0, 0, 10, true},
		{-1, 0, 10, false},
		{-1, 0, 11, false},
		{-1, 0, 12, false},
		{0, 0, 11, false},
		{1, 0, 10, false},
		{-1, 0, 12, false},
		{1, 0, 9, true},
		{0, 0, 9, true},
		{1, 0, 8, true},
	}
	for _, tc := range tcs {
		p := pts.P2{X: tc.x, Y: tc.y}
		q := p.Add(pts.P2{X: tc.w, Y: 0})
		h := NewHSeg(p, q)
		got := acceptableHSeg(h, hs)
		a.Equal(tc.expected, got, fmt.Sprintf("%v", tc))
	}
}

func TestAcceptableVSegs(t *testing.T) {
	a := assert.New(t)
	_, vs := baseSegs()
	tcs := []struct {
		x        int
		y        int
		h        int
		expected bool
	}{
		{0, 5, 5, true},
		{0, 4, 5, false},
		{0, 4, 6, false},
		{0, 4, 7, false},
		{0, 4, 7, false},
		{0, 5, 6, false},
		{0, 6, 5, false},
		{0, 6, 5, false},
		{0, 6, 4, true},
		{0, 6, 3, true},
	}
	for _, tc := range tcs {
		p := pts.P2{X: tc.x, Y: tc.y}
		q := p.Add(pts.P2{X: 0, Y: tc.h})
		v := NewVSeg(p, q)
		got := acceptableVSeg(v, vs)
		a.Equal(tc.expected, got, fmt.Sprintf("%v", tc))
	}
}
