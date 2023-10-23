package bitfield

type Bitfield []byte

func New(size int) Bitfield {
	bf := make([]byte, size/8+1)
	return bf
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
