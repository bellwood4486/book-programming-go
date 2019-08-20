package bitdiff

func CountBitDiffSHA256(c1, c2 [32]byte) int {
	return countBitDiff(c1[:], c2[:])
}

func countBitDiff(c1, c2 []byte) int {
	// 参考：https://github.com/torbiak/gopl/blob/master/ex4.1/shadiff.go
	count := 0
	for i := 0; i < len(c1) || i < len(c2); i++ {
		switch {
		case i >= len(c1):
			count += popCountBit(c2[i])
		case i >= len(c2):
			count += popCountBit(c1[i])
		default:
			count += popCountBit(c1[i] ^ c2[i])
		}
	}
	return count
}

func popCountBit(x byte) int {
	c := 0
	for ; x > 0; x >>= 1 {
		c += int(x & 0x01)
	}
	return c
}
