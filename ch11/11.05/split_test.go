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
		{"abc", ":", 1},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			words := strings.Split(tt.input, tt.sep)
			if got := len(words); got != tt.want {
				t.Errorf("strings.Split(%q, %q) returned %d words, want %d", tt.input, tt.sep, got, tt.want)
			}
		})
	}
}
