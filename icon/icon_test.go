package icon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinear(t *testing.T) {
	a := assert.New(t)
	tcs := []struct {
		name     string
		terms    []Term
		expected []Term
	}{
		{"id", []Term{{1, "x"}, {2, "y"}}, []Term{{1, "x"}, {2, "y"}}},
		{"sort", []Term{{2, "y"}, {1, "x"}}, []Term{{1, "x"}, {2, "y"}}},
		{"coalesce", []Term{{2, "x"}, {1, "x"}}, []Term{{3, "x"}}},
		{"zero", []Term{{-1, "x"}, {1, "x"}}, []Term{{0, "x"}}},
	}
	for _, tc := range tcs {
		got := NewLinear(tc.terms).Terms()
		a.Equal(tc.expected, got, tc.name)
	}
}
