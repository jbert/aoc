package bitfield

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	a := assert.New(t)

	bf := New(8)
	a.Equal(false, bf.Get(0))
	a.Equal(false, bf.Get(1))
	a.Equal(false, bf.Get(3))
	a.Equal(false, bf.Get(7))

	bf.Set(0)
	a.Equal(true, bf.Get(0))
	a.Equal(false, bf.Get(1))
	a.Equal(false, bf.Get(3))
	a.Equal(false, bf.Get(7))

	bf.Set(7)
	a.Equal(true, bf.Get(0))
	a.Equal(false, bf.Get(1))
	a.Equal(false, bf.Get(3))
	a.Equal(true, bf.Get(7))

	bf.Clear(0)
	a.Equal(false, bf.Get(0))
	a.Equal(false, bf.Get(1))
	a.Equal(false, bf.Get(3))
	a.Equal(true, bf.Get(7))

	bf.Clear(7)
	a.Equal(false, bf.Get(0))
	a.Equal(false, bf.Get(1))
	a.Equal(false, bf.Get(3))
	a.Equal(false, bf.Get(7))
}
