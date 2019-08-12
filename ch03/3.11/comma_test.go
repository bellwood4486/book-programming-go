package comma

import "testing"

func Test_comma(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want string
	}{
		{args{""}, ""},
		{args{"123"}, "123"},
		{args{"1234"}, "1,234"},
		{args{"1234567"}, "1,234,567"},
		{args{"-1.1234"}, "-1.1234"},
		{args{"+1234567"}, "+1,234,567"},
		{args{"-1234567.123"}, "-1,234,567.123"},
	}
	for _, tt := range tests {
		t.Run(tt.args.s, func(t *testing.T) {
			if got := comma(tt.args.s); got != tt.want {
				t.Errorf("comma() = %v, want %v", got, tt.want)
			}
		})
	}
}
