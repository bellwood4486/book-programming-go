package comma

import "bytes"

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	a := n % 3
	var buf bytes.Buffer
	buf.WriteString(s[:a])
	for i := a; i < n; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}

func commaRecursive(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
