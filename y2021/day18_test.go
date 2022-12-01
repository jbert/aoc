package y2021

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	a := assert.New(t)
	testCases := []struct {
		line     string
		expected string
	}{

		{"[1,10]", "[1,[5,5]]"},
		{"[1,11]", "[1,[5,6]]"},
	}
	for _, tc := range testCases {
		snum := lineToSnum(tc.line)
		didSplit := snum.checkDoSplit()
		a.True(didSplit, "these test cases should all split")
		a.Equal(tc.expected, snum.String(), fmt.Sprintf("Test split of %s", tc.line))
	}
}

func TestExplode(t *testing.T) {
	a := assert.New(t)
	testCases := []struct {
		line     string
		expected string
	}{

		{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
		{"[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"},
		{"[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"},
		{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
		{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
	}
	for _, tc := range testCases {
		snum := lineToSnum(tc.line)
		didExplode := snum.checkDoExplode()
		a.True(didExplode, "these test cases should all explode")
		a.Equal(tc.expected, snum.String(), fmt.Sprintf("Test explode of %s", tc.line))
	}
}

func TestMagnitude(t *testing.T) {
	a := assert.New(t)
	testCases := []struct {
		line     string
		expected int
	}{
		{"[[1,2],[[3,4],5]]", 143},
		{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
		{"[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
		{"[[[[3,0],[5,3]],[4,4]],[5,5]]", 791},
		{"[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
		{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
	}
	for _, tc := range testCases {
		snum := lineToSnum(tc.line)
		got := snum.Magnitude()
		a.Equal(tc.expected, got, fmt.Sprintf("Test magnitude of %s", tc.line))
	}
}
