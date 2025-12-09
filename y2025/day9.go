package y2025

import (
	"fmt"
	"io"
	"slices"

	"github.com/jbert/aoc/pts"
	"github.com/jbert/aoc/year"
	"github.com/jbert/fun"
)

type Day9 struct{ year.Year }

type HSeg struct {
	l pts.P2
	w int
}

func NewHSeg(p pts.P2, q pts.P2) HSeg {
	if p.Y != q.Y {
		panic("bad hseg")
	}
	l := min(p.X, q.X)
	r := max(p.X, q.X)
	return HSeg{l: pts.P2{X: l, Y: p.Y}, w: r - l + 1}
}

func acceptableHSeg(h HSeg, hs []HSeg) bool {
	for _, g := range hs {
		if g.Contains(h) {
			return true
		}
	}
	return false
}

func mergeHSegs(hs []HSeg) []HSeg {
	slices.SortFunc(hs, cmpHSeg)
	var mhs []HSeg
	for i, g := range hs {
		if i == 0 {
			continue
		}
		h := hs[i-1]
		if h.ContainsPt(g.l) {
			newHS = NewHSeg(h.l, g.l.Add(pts.P2{g.w, 0}))
			mhs = append(mhs, newHS)
		} else {
			mhs = append(mhs, h)
		}
	}
}

func (h HSeg) ContainsPt(p pts.P2) bool {
	return p.Y == h.l.Y && p.X >= h.l.X && p.X <= h.l.X+h.w
}

func (h HSeg) Contains(g HSeg) bool {
	return g.l.Y == h.l.Y && g.l.X >= h.l.X && g.l.X+g.w <= h.l.X+h.w
}

func cmpHSeg(a HSeg, b HSeg) int {
	c := pts.CmpP2(a.l, b.l)
	if c != 0 {
		return c
	}
	return a.w - b.w
}

type VSeg struct {
	b pts.P2
	h int
}

func NewVSeg(p pts.P2, q pts.P2) VSeg {
	if p.X != q.X {
		panic("bad vseg")
	}
	b := min(p.Y, q.Y)
	t := max(p.Y, q.Y)
	return VSeg{b: pts.P2{X: p.X, Y: b}, h: t - b + 1}
}

func (v VSeg) ContainsPt(p pts.P2) bool {
	return p.X == v.b.X && p.Y >= v.b.Y && p.Y <= v.b.Y+v.h
}

func (v VSeg) Contains(w VSeg) bool {
	return w.b.X == v.b.X && w.b.Y >= v.b.Y && w.b.Y+w.h <= v.b.Y+v.h
}

func cmpVSeg(a VSeg, b VSeg) int {
	c := pts.CmpP2(a.b, b.b)
	if c != 0 {
		return c
	}
	return a.h - b.h
}

func acceptableVSeg(v VSeg, vs []VSeg) bool {
	for _, w := range vs {
		if w.Contains(v) {
			return true
		}
	}
	return false
}

func acceptableRect(r pts.Rect, hs []HSeg, vs []VSeg) bool {
	l := NewVSeg(r.BL(), r.TL())
	rr := NewVSeg(r.BR(), r.TR())
	if !acceptableVSeg(l, vs) || !acceptableVSeg(rr, vs) {
		return false
	}
	b := NewHSeg(r.BL(), r.BR())
	t := NewHSeg(r.TL(), r.TR())
	if !acceptableHSeg(b, hs) || !acceptableHSeg(t, hs) {
		return false
	}
	return true
}

func (d *Day9) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")
	ps := fun.Map(pts.P2FromString, lines)
	fmt.Printf("%v\n", ps)
	slices.SortFunc(ps, pts.CmpP2)

	mxArea := 0
	// var mxRect pts.Rect
	for i := range ps {
		for j := range ps {
			if j <= i {
				continue
			}
			r := pts.NewRect(ps[i], ps[j])
			ra := r.Area()
			// fmt.Printf("RA %d R: %s\n", ra, r)
			if ra > mxArea {
				mxArea = ra
				// mxRect = r
			}
		}
	}
	// fmt.Printf("mx %d mxRect: %v\n", mxArea, mxRect)
	fmt.Printf("Part 1: %d\n", mxArea)

	// Unsort...
	ps = fun.Map(pts.P2FromString, lines)
	var hsegs []HSeg
	var vsegs []VSeg
	for i, q := range ps {
		var p pts.P2
		if i == 0 {
			p = ps[len(ps)-1]
		} else {
			p = ps[i-1]
		}
		if p.X == q.X {
			vsegs = append(vsegs, NewVSeg(p, q))
		} else if p.Y == q.Y {
			hsegs = append(hsegs, NewHSeg(p, q))
		} else {
			panic("bad input")
		}
		fmt.Printf("%s\n", p)
	}
	// TODO: need to merge adjacent
	fmt.Printf("vsegs: %v\n", vsegs)
	fmt.Printf("hsegs: %v\n", hsegs)

	mxArea = 0
	// var mxRect pts.Rect
	for i := range ps {
		for j := range ps {
			if j <= i {
				continue
			}
			r := pts.NewRect(ps[i], ps[j])
			if !acceptableRect(r, hsegs, vsegs) {
				continue
			}
			ra := r.Area()
			// fmt.Printf("RA %d R: %s\n", ra, r)
			if ra > mxArea {
				mxArea = ra
				// mxRect = r
			}
		}
	}

	// fmt.Printf("mx %d mxRect: %v\n", mxArea, mxRect)
	fmt.Printf("Part 2: %d\n", mxArea)

	return nil
}
