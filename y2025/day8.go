package y2025

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/jbert/aoc/num"
	"github.com/jbert/aoc/pts"
	"github.com/jbert/aoc/year"
	"github.com/jbert/fun"
	"github.com/jbert/set"
)

type Day8 struct{ year.Year }

func parsePoint(l string) pts.P3 {
	bits := strings.Split(l, ",")
	ns := fun.Map(num.MustAtoi, bits)
	return pts.P3{ns[0], ns[1], ns[2]}
}

type Edge struct {
	a pts.P3
	b pts.P3
}

func (e Edge) String() string {
	return fmt.Sprintf("%v -> %v", e.a, e.b)
}

func (e Edge) Length() float64 {
	return e.a.Sub(e.b).EuclideanLength()
}

type Circuits struct {
	cids map[pts.P3]int
}

func FromPts(ps []pts.P3) *Circuits {
	cid := 0
	cs := Circuits{cids: make(map[pts.P3]int)}
	for _, p := range ps {
		cs.cids[p] = cid
		cid++
	}
	return &cs
}

func (c *Circuits) join(e Edge) {
	// fmt.Printf("JOIN: %s\n", e)
	ca := c.cids[e.a]
	cb := c.cids[e.b]
	cid := min(ca, cb)
	for p := range c.cids {
		if c.cids[p] == ca || c.cids[p] == cb {
			c.cids[p] = cid
		}
	}
}

func (c *Circuits) maxCID() int {
	mxcid := 0
	for p := range c.cids {
		if c.cids[p] > mxcid {
			mxcid = c.cids[p]
		}
	}
	return mxcid
}

func (c *Circuits) CircuitLists() [][]pts.P3 {
	var circs [][]pts.P3
	mx := c.maxCID()
	for cid := range mx + 1 {
		var circ []pts.P3
		for p := range c.cids {
			if c.cids[p] == cid {
				circ = append(circ, p)
			}
		}
		if len(circ) > 0 {
			circs = append(circs, circ)
		}
	}
	return circs
}

func (d *Day8) Run(out io.Writer, lines []string) error {
	fmt.Fprintf(out, "Running\n")
	ps := fun.Map(parsePoint, lines)
	// fmt.Printf("ps %v\n", ps)
	distList := make(map[float64]set.Set[Edge])
	for i := range ps {
		for j := range ps {
			if j <= i {
				continue
			}
			e := Edge{ps[i], ps[j]}
			d := e.Length()

			s := distList[d]
			if s == nil {
				s = set.New[Edge]()
			}
			s.Insert(e)
			distList[d] = s
			// fmt.Printf("%f: %v\n", d, e)
		}
	}
	// fmt.Printf("%+v\n", distList)

	var ds []float64
	for d := range distList {
		ds = append(ds, d)
	}
	sort.Float64s(ds)
	// fmt.Printf("%+v\n", ds)
	// fmt.Printf("JB %v\n", distList[ds[0]])

	cs := FromPts(ps)
	// numConns := 10
	numConns := 1000
	done := 0
LOOPING:
	for _, d := range ds {
		es := distList[d].ToList()
		for _, e := range es {
			cs.join(e)
			done++
			if done == numConns {
				break LOOPING
			}
		}
	}
	circs := cs.CircuitLists()
	// for _, circ := range circs {
	// fmt.Printf("%d: cs %+v\n", len(circ), circ)
	// }
	sizes := fun.Map(func(l []pts.P3) int { return len(l) }, circs)
	sort.Ints(sizes)
	sizes = fun.Reverse(sizes)
	fmt.Printf("%v\n", sizes)

	fmt.Printf("Part 1: %d\n", sizes[0]*sizes[1]*sizes[2])
	// fmt.Printf("Part 2: %d\n", fun.Sum(joltages))

	return nil
}
