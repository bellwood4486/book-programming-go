package squash

import (
	"unicode"
	"unicode/utf8"
)

const space = ' '

func squash(s []byte) []byte {
	wIdx := 0
	var prev rune
	for _, r := range string(s) {
		if unicode.IsSpace(r) {
			if prev != space {
				s[wIdx] = space
				prev = space
				wIdx++
			}
			continue
		}

		buf := make([]byte, utf8.RuneLen(r))
		n := utf8.EncodeRune(buf, r)
		copy(s[wIdx:], buf)
		prev = r
		wIdx += n
	}
	return s[:wIdx]
}
