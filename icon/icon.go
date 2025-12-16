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

type Linear map[Var]Term

func NewLinear(terms []Term) Linear {
	m := make(map[Var]Term)
	for _, t := range terms {
		tt, ok := m[t.Var]
		if ok {
			t.Coeff += tt.Coeff
		}
		delete(m, t.Var)
		if t.Coeff != 0 {
			m[t.Var] = t
		}
	}
	return m
}

func (l Linear) Terms() []Term {
	ts := make([]Term, 0)
	for _, t := range l {
		ts = append(ts, t)
	}
	slices.SortFunc(ts, CmpTerm)
	return ts
}
