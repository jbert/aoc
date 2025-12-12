package matrix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMult(t *testing.T) {
	a := assert.New(t)
	id3 := NewFromRows([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}})
	diag2s := NewFromRows([][]int{{2, 0, 0}, {0, 2, 0}, {0, 0, 2}})
	diag4s := NewFromRows([][]int{{4, 0, 0}, {0, 4, 0}, {0, 0, 4}})
	tcs := []struct {
		a           Mat
		b           Mat
		expected    Mat
		expectedErr error
	}{
		{id3, id3, id3, nil},
		{id3, diag2s, diag2s, nil},
		{diag2s, diag2s, diag4s, nil},

		{NewFromRows([][]int{{1, 2, 3}, {4, 5, 6}}),
			NewFromRows([][]int{{1, 2}, {3, 4}, {5, 6}}),
			NewFromRows([][]int{{1 + 2*3 + 3*5, 1*2 + 2*4 + 3*6}, {4*1 + 5*3 + 6*5, 4*2 + 5*4 + 6*6}}), nil},
	}

	for i, tc := range tcs {
		fmt.Printf("%v\n", tc)
		got, err := Mult(tc.a, tc.b)
		a.Equal(err, tc.expectedErr, fmt.Sprintf("%d: correct error", i))
		a.Equal(got, tc.expected, fmt.Sprintf("%d: correct", i))
	}
}
