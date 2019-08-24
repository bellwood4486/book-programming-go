package rotate

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotateOriginal(s []int, shift int) {
	reverse(s[:shift])
	reverse(s[shift:])
	reverse(s)
}

func rotate(s []int, shift int) {
	tmp := make([]int, shift)
	copy(tmp, s[:shift])
	copy(s, s[shift:])
	copy(s[len(s)-shift:], tmp)
}
