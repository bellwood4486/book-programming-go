package wordfreq

import (
	"bufio"
	"io"
)

func wordfreq(rd io.Reader) map[string]int {
	freq := make(map[string]int)

	sc := bufio.NewScanner(rd)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		t := sc.Text()
		freq[t]++
	}
	return freq
}
