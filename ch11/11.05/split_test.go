package split

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		input string
		sep   string
		want  int
	}{
		{"a:b:c", ":", 3},
	}
	for _, test := range tests {
		words := strings.Split(test.input, test.sep)
		if got := len(words); got != test.want {
			t.Errorf("strings.Split(%q, %q) returned %d words, want %d", test.input, test.sep, got, test.want)
		}
	}
}
