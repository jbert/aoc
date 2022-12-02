package y2022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRPS(t *testing.T) {
	a := assert.New(t)
	a.True(R.Beats(S))
	a.True(S.Beats(P))
	a.True(P.Beats(R))
}
