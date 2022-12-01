package pts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP3FromString(t *testing.T) {
	a := assert.New(t)
	testCases := []struct {
		l string
		p P3
	}{
		{"0,0,0", P3{0, 0, 0}},
		{"1,0,0", P3{1, 0, 0}},
		{"0,1,0", P3{0, 1, 0}},
		{"0,0,1", P3{0, 0, 1}},
	}
	for _, tc := range testCases {
		p := P3FromString(tc.l)
		a.Equal(tc.p, p, "Can parse [%s] correctly to [%s]", tc.l, tc.p)
	}
}
