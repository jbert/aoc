package y2022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRPS(t *testing.T) {
	a := assert.New(t)
	a.True(R.WinsOver(S))
	a.True(S.WinsOver(P))
	a.True(P.WinsOver(R))
}
