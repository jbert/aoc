package icon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	a := assert.New(t)
	xPlusYEq0 := NewAffine([]Term{{1, "x"}, {1, "y"}}, 0)
	YEq1 := NewAffine([]Term{{1, "y"}}, 1)
	tcs := []struct {
		name        string
		constraints []Constraint
		hasSoln     bool
		expected    []Term
	}{
		{
			"hello world",
			[]Constraint{xPlusYEq0, YEq1},
			true,
			[]Term{{-1, "x"}, {1, "y"}},
		},
	}

	for _, tc := range tcs {
		soln, err := Solve(tc.constraints)
		if tc.hasSoln {
			a.Equal(tc.expected, soln, tc.name)
		} else {
			a.Equal(err, ErrNoSoln)
		}
	}
}
