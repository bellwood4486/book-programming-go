package charcount

import (
	"bufio"
	"github.com/pkg/errors"
	"io"
	"unicode"
	"unicode/utf8"
)

const utflenSize = utf8.UTFMax + 1
const (
	letter int = iota
	number
	other
)

type countMap map[rune]int
type categorizedCountMap map[int]countMap

func charcount(rd io.Reader) (counts categorizedCountMap, utflen [utflenSize]int, invalid int, err error) {
	counts = make(categorizedCountMap)
	in := bufio.NewReader(rd)
	for {
		r, size, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return counts, utflen, invalid, errors.Errorf("charcount: %v\n", err)
		}
		if r == unicode.ReplacementChar && size == 1 {
			invalid++
			continue
		}
		c := category(r)
		if counts[c] == nil {
			counts[c] = make(countMap)
		}
		counts[c][r]++
		utflen[size]++
	}
	return counts, utflen, invalid, nil
}

func category(r rune) int {
	if unicode.IsLetter(r) {
		return letter
	}
	if unicode.IsNumber(r) {
		return number
	}
	return other
}
