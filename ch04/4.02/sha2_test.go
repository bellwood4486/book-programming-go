package main

import (
	"testing"
)

func Test_sha2String(t *testing.T) {
	type args struct {
		in   []byte
		kind string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"sha256empty", args{[]byte{}, KindSHA256}, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", false},
		{"sha384empty", args{[]byte{}, KindSHA384}, "38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b", false},
		{"sha512empty", args{[]byte{}, KindSHA512}, "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e", false},
		{"unknownKind", args{[]byte{}, "unknown"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sha2String(tt.args.in, tt.args.kind)
			if (err != nil) != tt.wantErr {
				t.Errorf("sha2String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sha2String() got = %v, want %v", got, tt.want)
			}
		})
	}
}
