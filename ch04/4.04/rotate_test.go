package rotate

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		s     []int
		shift int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"2shift", args{[]int{0, 1, 2, 3, 4}, 2}, []int{2, 3, 4, 0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.s, tt.args.shift)
			if !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("rotate() = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}
