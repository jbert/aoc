package y2025

import (
	"fmt"
	"io"
	"maps"
	"slices"
	"strings"

	"github.com/jbert/aoc/grid"
	"github.com/jbert/aoc/pts"
	"github.com/jbert/aoc/year"
	"github.com/jbert/fun"
)

type Day9 struct{ year.Year }

type Seg struct {
	fr pts.P2
	to pts.P2
}

func iAbs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func (s Seg) Dir() pts.P2 {
	dir := s.to.Sub(s.fr)
	if dir.X != 0 {
		dir.X /= iAbs(dir.X)
	}
	if dir.Y != 0 {
		dir.Y /= iAbs(dir.Y)
	}
	return dir
}

func (s Seg) Pts() []pts.P2 {
	dir := s.Dir()
	var ps []pts.P2
	for p := s.fr; ; p = p.Add(dir) {
		ps = append(ps, p)
		if p.Equals(s.to) {
			break
		}
	}
	return ps
}

func (d *Day9) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")
	ps := fun.Map(pts.P2FromString, lines)
	fmt.Printf("%v\n", ps)
	slices.SortFunc(ps, pts.CmpP2)

	mxArea := 0
	mxx := 0
	mxy := 0
	// var mxRect pts.Rect
	for i := range ps {
		if ps[i].X > mxx {
			mxx = ps[i].X
		}
		if ps[i].Y > mxy {
			mxy = ps[i].Y
		}
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

	// Unsort, so we can loop end->start
	ps = fun.Map(pts.P2FromString, lines)
	// sps := set.NewFromList(ps)
	// if sps.Size() != len(ps) {
	// panic("points listed twice")
	// }
	// isRed := func(p pts.P2) bool { return sps.Contains(p) }

	g := grid.New[bool](mxx+1, mxy+1)
	p2seg := make(map[pts.P2]Seg)

	for i, q := range ps {
		var p pts.P2
		if i == 0 {
			p = ps[len(ps)-1]
		} else {
			p = ps[i-1]
		}
		seg := Seg{fr: p, to: q}
		p2seg[q] = seg
	}

	segs := maps.Values(p2seg)

	for seg := range segs {
		ps := seg.Pts()
		for _, p := range ps {
			g.SetPt(p, true)
		}
	}

	printGrid(out, g)
	// fmt.Printf("Part 2: %d\n", mxArea)
	return nil
}

func printGrid(w io.Writer, g grid.Grid[bool]) {
	for _, row := range g.Rows() {
		bits := fun.Map(func(isOK bool) string {
			if isOK {
				return "X"
			} else {
				return "."
			}
		}, row)
		fmt.Fprintf(w, "%s\n", strings.Join(bits, ""))
	}
	fmt.Fprintf(w, "\n")
}
