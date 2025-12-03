package y2025

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumToID(t *testing.T) {
	a := assert.New(t)
	testCases := []struct {
		n        int
		copies   int
		expected int
	}{
		{123, 1, 123},
		{123, 2, 123123},
		{1, 6, 111111},
		{12, 3, 121212},
	}
	for _, tc := range testCases {
		got := numToID(tc.n, tc.copies)
		a.Equal(tc.expected, got, fmt.Sprintf("r %d", tc.n))
	}
}

func TestFirstNDigits(t *testing.T) {
	a := assert.New(t)
	testCases := []struct {
		n        int
		nDig     int
		expected int
	}{
		{123456, 1, 1},
		{123456, 2, 12},
		{123456, 3, 123},
		{123456, 4, 1234},
		{123456, 5, 12345},
	}
	for _, tc := range testCases {
		got := firstNDigits(tc.n, tc.nDig)
		a.Equal(tc.expected, got, fmt.Sprintf("r %d", tc.n))
	}
}

func TestSameDigitRanges(t *testing.T) {
	a := assert.New(t)
	testCases := []struct {
		r        Range
		expected []Range
	}{
		{Range{1, 10}, []Range{{1, 9}, {10, 10}}},
		{Range{10, 20}, []Range{{10, 20}}},
		{Range{10, 100}, []Range{{10, 99}, {100, 100}}},
	}
	for _, tc := range testCases {
		got := tc.r.sameDigitRanges()
		slices.SortFunc(got, cmpRange)
		slices.SortFunc(tc.expected, cmpRange)
		a.Equal(tc.expected, got, fmt.Sprintf("r %s", tc.r))
	}
}

func TestDownToEven(t *testing.T) {
	a := assert.New(t)
	testCases := []struct {
		n        int
		expected int
	}{
		{115, 99},
	}
	for _, tc := range testCases {
		got := downToEvenDigits(tc.n)
		a.Equal(tc.expected, got, fmt.Sprintf("n %d", tc.n))
	}
}

func TestSecondHalf(t *testing.T) {
	a := assert.New(t)
	testCases := []struct {
		n        int
		expected int
	}{
		{123456, 456},
		{1256, 56},
		{11, 1},
		{22, 2},
	}
	for _, tc := range testCases {
		got := secondHalf(tc.n)
		a.Equal(tc.expected, got, fmt.Sprintf("n %d", tc.n))
	}
}

func TestFirstHalf(t *testing.T) {
	a := assert.New(t)
	testCases := []struct {
		n        int
		expected int
	}{
		{123456, 123},
		{1256, 12},
		{11, 1},
		{22, 2},
	}
	for _, tc := range testCases {
		got := firstHalf(tc.n)
		a.Equal(tc.expected, got, fmt.Sprintf("n %d", tc.n))
	}
}

func TestPow10(t *testing.T) {
	a := assert.New(t)
	testCases := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 10},
		{3, 100},
		{4, 1000},
	}
	for _, tc := range testCases {
		got := pow10(tc.n)
		a.Equal(tc.expected, got, fmt.Sprintf("n %d", tc.n))
	}
}

func TestNumDigits(t *testing.T) {
	a := assert.New(t)
	testCases := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{9, 1},
		{333, 3},
		{33, 2},
		{5555, 4},
	}
	for _, tc := range testCases {
		got := numDigits(tc.n)
		a.Equal(tc.expected, got, fmt.Sprintf("n %d", tc.n))
	}
}

func TestUpToEvenDigits(t *testing.T) {
	a := assert.New(t)
	testCases := []struct {
		n        int
		expected int
	}{
		{1, 10},
		{9, 10},
		{333, 1000},
		{33, 33},
		{5555, 5555},
	}
	for _, tc := range testCases {
		got := upToEvenDigits(tc.n)
		a.Equal(tc.expected, got, fmt.Sprintf("n %d", tc.n))
	}
}
