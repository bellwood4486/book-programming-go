package rev

import (
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		s *[5]int
	}
	tests := []struct {
		name string
		args args
		want [5]int
	}{
		{"0to4", args{&[5]int{0, 1, 2, 3, 4}}, [5]int{4, 3, 2, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverse(tt.args.s)
			if !reflect.DeepEqual(*tt.args.s, tt.want) {
				t.Errorf("reverse() = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}

func Test_reverseSlice(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"0to4", args{[]int{0, 1, 2, 3, 4}}, []int{4, 3, 2, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseSlice(tt.args.s)
			if !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("reverse() = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}
