package icon

import (
	"errors"
	"fmt"
	"maps"
	"math"
	"slices"
	"sort"

	"github.com/jbert/set"
)

type Term struct {
	Coeff int
	Var   string
}

func CmpTerm(a Term, b Term) int {
	if a.Var < b.Var {
		return -1
	} else if a.Var > b.Var {
		return +1
	}
	return a.Coeff - b.Coeff
}

type Vec map[string]int

type Affine struct {
	Vec   Vec
	Const int
}

func (a Affine) Range(v string) Range {
	vs := set.NewFromList(a.Vars())
	if !vs.Contains(v) {
		return Range{0, 0}
	}
	// Could do better here :-)
	// But would need to pass in more info
	m, ok := a.Vec[v]
	if !ok {
		panic(fmt.Sprintf("checked we have var but not present in vec"))
	}
	c := a.Const / m
	return Range{-c, c}
}

func (a Affine) Refine(vec Vec) (Constraint, error) {
	x := 0
	var refined Affine
	for _, v := range a.Vars() {
		m := a.Vec[v]
		if vec.HasVar(v) {
			x += m * vec[v]
		} else {
			refined.Vec.AddVar(v, m)
		}
	}
	refined.Const = a.Const - x
	return refined, nil
}

type Constraint interface {
	Vars() []string
	Refine(v Vec) (Constraint, error)
	Range(v string) Range
}

// Half empty, lo <= x < hi
// We steal MaxInt, MinInt for 'infinity'
type Range struct {
	lo int
	hi int
}

func (r Range) Width() int {
	if r.hi == math.MaxInt || r.lo == math.MinInt {
		return math.MaxInt
	}
	return r.hi - r.lo
}

func (r Range) Intersect(s Range) Range {
	if r.lo > s.lo {
		return s.Intersect(r)
	}
	if r.hi >= s.lo {
		return Range{0, 0}
	}
	return Range{s.lo, r.hi}
}

func NewAffine(terms []Term, c int) Affine {
	return Affine{Vec: NewVec(terms), Const: c}
}

func (a Affine) Vars() []string {
	return a.Vec.Vars()
}

type EqZero Affine

type LtZero Affine

func NewVec(terms []Term) Vec {
	m := make(map[string]int)
	for _, t := range terms {
		_, ok := m[t.Var]
		if ok {
			panic("repeated terms in vec")
		}
		m[t.Var] = t.Coeff
	}
	return m
}

func (v Vec) Terms() []Term {
	ts := make([]Term, 0)
	for v, c := range v {
		ts = append(ts, Term{Coeff: c, Var: v})
	}
	slices.SortFunc(ts, CmpTerm)
	return ts
}

func (v Vec) Vars() []string {
	vsi := maps.Keys(v)
	vs := slices.Collect(vsi)
	vs = sort.StringSlice(vs)
	return vs
}

func (v Vec) HasVar(s string) bool {
	_, ok := v[s]
	return ok
}

func (v Vec) Copy() Vec {
	vv := make(map[string]int)
	for s, m := range v {
		vv[s] = m
	}
	return vv
}

func (v Vec) AddVar(s string, m int) Vec {
	vv := v.Copy()
	vv[s] = m
	return vv
}

func (v Vec) SameVars(u Vec) ([]string, bool) {
	uvs := u.Vars()
	vvs := v.Vars()
	if len(uvs) != len(vvs) {
		return nil, false
	}
	for i := range uvs {
		if uvs[i] != vvs[i] {
			return nil, false
		}
	}
	return vvs, true
}

func (v Vec) Dot(u Vec) int {
	vs, same := v.SameVars(u)
	if !same {
		panic("variable mismatch")
	}
	d := 0
	for _, vv := range vs {
		d += v[vv] * u[vv]
	}
	return d
}

var ErrNoSoln = errors.New("no solution to constraints")

func Solve(cs []Constraint) (Vec, error) {
	v := NewVec([]Term{})
	return solveVec(cs, v)
}

func solveVec(cs []Constraint, v Vec) (Vec, error) {
	var freeVars set.Set[string]
	for _, c := range cs {
		cc, err := c.Refine(v)
		if err != nil {
			return nil, err
		}
		freeVars = freeVars.Union(set.NewFromList(cc.Vars()))
	}
	// for each var, get a range from each constraint, then loop
	freeVars.ForEach(func(v string) {
		var vr Range
		for i, c := range cs {
			cr := c.Range(v)
			if i == 0 {
				vr = cr
			} else {
				vr = vr.Intersect(cr)
			}
		}
		...loop over range and recurse
	})
	return NewVec([]Term{{1, "x"}}), nil
}
