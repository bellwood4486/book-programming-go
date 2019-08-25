package squash

import (
	"reflect"
	"testing"
)

var spaceRunes = []rune{'\t', '\n', '\v', '\f', '\r', ' ', '\u0085', '\u00a0'}

func Test_squash(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"firstSpace", args{toBytes(spaceRunes, []rune("ab世界"))}, []byte(" ab世界")},
		{"midSpace", args{toBytes([]rune("a世"), spaceRunes, []rune("b界"))}, []byte("a世 b界")},
		{"lastSpace", args{toBytes([]rune("ab世界"), spaceRunes)}, []byte("ab世界 ")},
		{"spaceOnly", args{toBytes(spaceRunes)}, []byte(" ")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squash(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("squash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func toBytes(vals ...[]rune) []byte {
	var c []rune
	for _, val := range vals {
		c = append(c, val...)
	}
	return []byte(string(c))
}
