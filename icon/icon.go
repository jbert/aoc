package icon

import "slices"

type Var string

type Term struct {
	Coeff int
	Var   Var
}

func CmpTerm(a Term, b Term) int {
	if a.Var < b.Var {
		return -1
	} else if a.Var > b.Var {
		return +1
	}
	return a.Coeff - b.Coeff
}

type Vec map[Var]int

// type (v Vec) Dot(u Vec) int {
// }

type Linear Vec

func NewLinear(terms []Term) Linear {
	m := make(map[Var]int)
	for _, t := range terms {
		c, ok := m[t.Var]
		if ok {
			t.Coeff += c
		}
		m[t.Var] = t.Coeff
	}
	return m
}

func (l Linear) Terms() []Term {
	ts := make([]Term, 0)
	for v, c := range l {
		ts = append(ts, Term{Coeff: c, Var: v})
	}
	slices.SortFunc(ts, CmpTerm)
	return ts
}
