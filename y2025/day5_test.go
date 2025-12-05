package y2025

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeJoin(t *testing.T) {
	a := assert.New(t)
	tcs := []struct {
		a        Range
		b        Range
		expected *Range
	}{
		{Range{0, 5}, Range{5, 10}, &Range{0, 10}},
		{Range{5, 10}, Range{0, 5}, &Range{0, 10}},
		{Range{0, 10}, Range{0, 5}, &Range{0, 10}},
		{Range{0, 10}, Range{2, 5}, &Range{0, 10}},
		{Range{0, 10}, Range{5, 10}, &Range{0, 10}},
		{Range{0, 5}, Range{6, 10}, nil},
		{Range{6, 10}, Range{0, 5}, nil},
	}

	for _, tc := range tcs {
		r, err := tc.a.Join(tc.b)
		if err == nil {
			a.Equal(*tc.expected, r, fmt.Sprintf("%v", tc))
		} else {
			a.Nil(tc.expected, fmt.Sprintf("%v", tc))
		}
	}
}
