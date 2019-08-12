package comma

import (
	"strings"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	sign := ""
	if i := strings.IndexAny(s, "+-"); i != -1 {
		sign, s = s[:i], s[i:]
	}
	decimalPart := ""
	if i := strings.LastIndex(s, "."); i != -1 {
		decimalPart, s = s[i:], s[:i]
	}

	return sign + commaIntPart(s) + decimalPart
}

func commaIntPart(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaIntPart(s[:n-3]) + "," + s[n-3:]
}
