package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountTable(x uint64) int {
	c := 0
	for i := uint(0); i < 8; i++ {
		c += int(pc[byte(x>>(i*8))])
	}
	return c
}

func PopCountBit(x uint64) int {
	c := 0
	for ; x > 0; x >>= 1 {
		c += int(x & 0x01)
	}
	return c
}

func PopCountClear(x uint64) int {
	c := 0
	for ; x > 0; x = x & (x - 1) {
		c++
	}
	return c
}
