package charcount

import (
	"strings"
	"testing"
)

func TestCountRunes(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"", 0},
		{"a", 1},
		{"ab", 2},
		{"あ", 1},
		{"あい", 2},
	}

	for _, test := range tests {
		got, err := CountRunes(strings.NewReader(test.input))
		if err != nil {
			t.Errorf("CountRunes(%q): %v", test.input, err)
			continue
		}
		if got != test.want {
			t.Errorf("CountRunes(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}
