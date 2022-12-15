package y2022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay13Parse(t *testing.T) {
	a := assert.New(t)
	tcs := []string{
		"[1,1,3,1,1]",
		"[1,1,5,1,1]",
		"[[1],[2,3,4]]",
		"[[1],4]",
		"[9]",
		"[[8,7,6]]",
		"[[4,4],4,4]",
		"[[4,4],4,4,4]",
		"[7,7,7,7]",
		"[7,7,7]",
		"[]",
		"[3]",
		"[[[]]]",
		"[[]]",
		"[1,[2,[3,[4,[5,6,7]]]],8,9]",
		"[1,[2,[3,[4,[5,6,0]]]],8,9]",
	}
	for _, tc := range tcs {
		p := lineToPacket(tc)
		s := p.String()
		a.Equal(tc, s, tc)
	}
}
