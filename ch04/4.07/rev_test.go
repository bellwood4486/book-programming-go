package rev

import (
	"reflect"
	"testing"
)

func Test_reverseUTF8(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"mix", args{[]byte("世界abc")}, []byte("cba界世")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseUTF8(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseUTF8() = %v, want %v", got, tt.want)
			}
		})
	}
}
