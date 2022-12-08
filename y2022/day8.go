package y2022

import (
	"fmt"
	"io"

	"github.com/jbert/aoc"
	"github.com/jbert/aoc/grid"
)

type Day8 struct{ Year }

func NewDay8() *Day8 {
	d := Day8{}
	return &d
}

func (d *Day8) Run(out io.Writer, lines []string) error {
	g := aoc.IntGrid(lines)
	w := g.Width()
	h := g.Height()

	// Do left and right
	lVis := grid.New[bool](w, h)
	rVis := grid.New[bool](w, h)
	for j := 0; j < h; j++ {
		lMaxHeight := -1
		rMaxHeight := -1
		for lCol := 0; lCol < w; lCol++ {
			rCol := w - 1 - lCol

			lHeight := g.Get(lCol, j)
			rHeight := g.Get(rCol, j)

			if rHeight > rMaxHeight {
				rVis.Set(rCol, j, true)
				rMaxHeight = rHeight
			}

			if lHeight > lMaxHeight {
				lVis.Set(lCol, j, true)
				lMaxHeight = lHeight
			}
		}
	}

	// Do up and down
	uVis := grid.New[bool](w, h)
	dVis := grid.New[bool](w, h)
	for i := 0; i < w; i++ {
		uMaxHeight := -1
		dMaxHeight := -1
		for uCol := 0; uCol < h; uCol++ {
			dCol := h - 1 - uCol

			uHeight := g.Get(i, uCol)
			dHeight := g.Get(i, dCol)

			if uHeight > uMaxHeight {
				uVis.Set(i, uCol, true)
				uMaxHeight = uHeight
			}

			if dHeight > dMaxHeight {
				dVis.Set(i, dCol, true)
				dMaxHeight = dHeight
			}
		}
	}

	fmt.Printf("%v\n", lVis)
	fmt.Printf("%v\n", rVis)
	fmt.Printf("%v\n", uVis)
	fmt.Printf("%v\n", dVis)

	or := func(a, b bool) bool { return a || b }
	vis := lVis.Combine(rVis, or)
	vis = vis.Combine(lVis, or)
	vis = vis.Combine(uVis, or)
	vis = vis.Combine(dVis, or)

	fmt.Printf("VIS:\n%v\n", vis)
	numVisible := 0
	vis.ForEachVal(func(v bool) {
		if v {
			numVisible++
		}
	})
	fmt.Printf("Part 1: %d\n", numVisible)

	return nil
}
