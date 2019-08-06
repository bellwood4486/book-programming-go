package charcount

import (
	"bufio"
	"io"
)

// CountRunes はUnicodeの文字数を数えます。
func CountRunes(reader io.Reader) (int, error) {
	br := bufio.NewReader(reader)
	counts := 0
	for {
		_, _, err := br.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
		counts++
	}
	return counts, nil
}
