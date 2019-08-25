package wordfreq

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_wordfreq(t *testing.T) {
	type args struct {
		rd io.Reader
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{"3words", args{strings.NewReader("dog cat cat")},
			map[string]int{"dog": 1, "cat": 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordfreq(tt.args.rd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordfreq() = %v, want %v", got, tt.want)
			}
		})
	}
}
