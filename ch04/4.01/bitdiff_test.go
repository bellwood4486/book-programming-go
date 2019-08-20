package bitdiff

import (
	"testing"
)

func Test_countBitDiff(t *testing.T) {
	type args struct {
		c1 []byte
		c2 []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"same digit", args{[]byte{0}, []byte{1}}, 1},
		{"different digit", args{[]byte{0}, []byte{1, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBitDiff(tt.args.c1, tt.args.c2); got != tt.want {
				t.Errorf("countBitDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountBitDiffSHA256(t *testing.T) {
	type args struct {
		c1 [32]byte
		c2 [32]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"same", args{pad0(1), pad0(1)}, 0},
		{"different", args{pad0(1, 1), pad0(1)}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountBitDiffSHA256(tt.args.c1, tt.args.c2); got != tt.want {
				t.Errorf("CountBitDiffSHA256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func pad0(elems ...byte) [32]byte {
	ret := [32]byte{}
	for i, v := range elems {
		ret[i] = v
	}
	return ret
}
