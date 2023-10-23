package bitfield

type Bitfield []byte

func New(size int) Bitfield {
	bf := make([]byte, size/8+1)
	return bf
}

func (bf Bitfield) Len() int {
	return len(bf) * 8
}

func (bf Bitfield) Set(n int) {
	ibyte := n / 8
	ibit := n % 8
	mask := byte(0x1) << ibit
	bf[ibyte] |= mask
}

func (bf Bitfield) Get(n int) bool {
	ibyte := n / 8
	ibit := n % 8
	mask := byte(0x1) << ibit
	return bf[ibyte]&mask > 0
}

func (bf Bitfield) Clear(n int) {
	ibyte := n / 8
	ibit := n % 8
	mask := ^byte(0x1 << ibit)
	bf[ibyte] &= mask
}

type B64 uint64

func (b *B64) Set(n int8) {
	mask := B64(0x1) << n
	*b |= mask
}

func (b B64) Get(n int8) bool {
	mask := B64(0x1) << n
	return b&mask > 0
}

func (b *B64) Clear(n int8) {
	mask := ^B64(0x1 << n)
	*b &= mask
}
