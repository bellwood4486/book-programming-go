package uniq

func uniq(s []string) []string {
	if len(s) == 0 {
		return s
	}

	idx := 0
	for _, v := range s {
		if s[idx] != v {
			idx++
			s[idx] = v
		}
	}
	return s[:idx+1]
}
