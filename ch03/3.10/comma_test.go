package comma

import "testing"

func Test_comma(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0digit", args{""}, ""},
		{"3digit", args{"123"}, "123"},
		{"4digit", args{"1234"}, "1,234"},
		{"7digit", args{"1234567"}, "1,234,567"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := comma(tt.args.s); got != tt.want {
				t.Errorf("comma() = %v, want %v", got, tt.want)
			}
		})
	}
}
