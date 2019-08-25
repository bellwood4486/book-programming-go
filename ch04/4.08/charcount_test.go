package charcount

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_charcount(t *testing.T) {
	type args struct {
		rd io.Reader
	}
	tests := []struct {
		name        string
		args        args
		wantCounts  categorizedCountMap
		wantUtflen  [utflenSize]int
		wantInvalid int
		wantErr     bool
	}{
		{"mix", args{strings.NewReader("aa世界")},
			categorizedCountMap{letter: {'a': 2, '世': 1, '界': 1}},
			[utflenSize]int{0, 2, 0, 2, 0}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCounts, gotUtflen, gotInvalid, err := charcount(tt.args.rd)
			if (err != nil) != tt.wantErr {
				t.Errorf("charcount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCounts, tt.wantCounts) {
				t.Errorf("charcount() gotCounts = %v, want %v", gotCounts, tt.wantCounts)
			}
			if !reflect.DeepEqual(gotUtflen, tt.wantUtflen) {
				t.Errorf("charcount() gotUtflen = %v, want %v", gotUtflen, tt.wantUtflen)
			}
			if gotInvalid != tt.wantInvalid {
				t.Errorf("charcount() gotInvalid = %v, want %v", gotInvalid, tt.wantInvalid)
			}
		})
	}
}
