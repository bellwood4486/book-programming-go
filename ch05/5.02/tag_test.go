package tag

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_tagfreq(t *testing.T) {
	type args struct {
		rd io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]int
		wantErr bool
	}{
		{"sample",
			// language=HTML
			args{strings.NewReader(`<html><head></head><body><p>test1</p><p>test2</p></body></html>`)},
			map[string]int{"html": 1, "head": 1, "body": 1, "p": 2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tagfreq(tt.args.rd)
			if (err != nil) != tt.wantErr {
				t.Errorf("tagfreq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tagfreq() got = %v, want %v", got, tt.want)
			}
		})
	}
}
