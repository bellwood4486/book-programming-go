package uniq

import (
	"reflect"
	"testing"
)

func Test_uniq(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"dupFirst", args{[]string{"a", "a", "b", "c"}}, []string{"a", "b", "c"}},
		{"dupMiddle", args{[]string{"a", "b", "b", "c"}}, []string{"a", "b", "c"}},
		{"dupLast", args{[]string{"a", "b", "c", "c"}}, []string{"a", "b", "c"}},
		{"empty", args{[]string{}}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniq(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uniq() = %v, want %v", got, tt.want)
			}
		})
	}
}
