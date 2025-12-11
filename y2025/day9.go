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
	"github.com/jbert/set"
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
	return s.to.Sub(s.fr).Normalise()
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

func (s Seg) isVertical() bool {
	return s.fr.X == s.to.X
}

func (s Seg) reverse() Seg {
	return Seg{fr: s.to, to: s.fr}
}

func (s Seg) isLeftTurn(t Seg) bool {
	sDir := s.Dir()
	tDir := t.Dir()
	if sDir.Equals(tDir) {
		panic("logic error: same dir")
	}
	if sDir.Equals(tDir.Neg()) {
		panic("logic error: reverse dir")
	}
	if sDir.Equals(pts.N) {
		return sDir.Equals(pts.W)
	} else if tDir.Equals(pts.E) {
		return tDir.Equals(pts.N)
	} else if sDir.Equals(pts.S) {
		return tDir.Equals(pts.E)
	} else if sDir.Equals(pts.W) {
		return tDir.Equals(pts.S)
	} else {
		panic(fmt.Sprintf("wtf: %s => %s", sDir, tDir))
	}
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
	p2segFr := make(map[pts.P2]Seg)
	p2segTo := make(map[pts.P2]Seg)

	for i, q := range ps {
		var p pts.P2
		if i == 0 {
			p = ps[len(ps)-1]
		} else {
			p = ps[i-1]
		}
		seg := Seg{fr: p, to: q}
		p2segFr[q] = seg
		p2segTo[p] = seg
	}

	// Draw in the border as true
	segs := maps.Values(p2segFr)
	for seg := range segs {
		ps := seg.Pts()
		for _, p := range ps {
			g.SetPt(p, true)
		}
	}
	fmt.Printf("JB1 - drawn border\n")

	// Flood fill the outside as true
	todo := set.NewFromList([]pts.P2{{X: 0, Y: 0}})
	imsg := 0
	mxtodo := g.Width() * g.Height()
ADDING:
	for todo.Size() > 0 {
		if imsg%100000 == 0 {
			fmt.Printf("%d %5f%% TODO: %d\n", imsg, 100*float64(imsg)/float64(mxtodo), todo.Size())
		}
		p, err := todo.Take()
		if err != nil {
			panic("can't take from non-empty set")
		}
		// Stop at border and also anything we may have already done
		if g.GetPt(p) {
			continue ADDING
		}

		g.SetPt(p, true)

		nxt := g.CardinalNeighbourPts(p)
		for _, q := range nxt {
			if !g.GetPt(q) {
				todo.Insert(q)
			}
		}
		imsg++
	}
	fmt.Printf("JB1 - done flood fill\n")

	// Invert grid
	g = grid.Fmap(g, func(b bool) bool { return !b })
	fmt.Printf("JB1 - done invert\n")
	// Re-draw in the border as true
	for seg := range segs {
		ps := seg.Pts()
		for _, p := range ps {
			g.SetPt(p, true)
		}
	}
	fmt.Printf("JB1 - re-drawn border\n")

	// Run same algorithm, but check area all true
	mxArea = 0
	mxx = 0
	mxy = 0
	for i := range ps {
	RECT:
		for j := range ps {
			if j <= i {
				continue
			}
			r := pts.NewRect(ps[i], ps[j])
			// fmt.Printf("JB - ps[i] %s ps[j] %s r %+v\n", ps[i], ps[j], r)
			ps := r.AllPts()
			// fmt.Printf("JB - g: %+v\n", g)
			for _, p := range ps {
				if !g.GetPt(p) {
					continue RECT
				}
			}
			ra := r.Area()

			// fmt.Printf("RA %d R: %s\n", ra, r)
			if ra > mxArea {
				mxArea = ra
				// mxRect = r
			}
		}
	}
	// printGrid(out, g)
	fmt.Printf("Part 2: %d\n", mxArea)
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
