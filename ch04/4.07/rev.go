package rev

import (
	"unicode/utf8"
)

// 参考：https://github.com/torbiak/gopl/blob/master/ex4.7/reverse.go
func reverseUTF8(s []byte) []byte {
	for i := 0; i < len(s); {
		_, size := utf8.DecodeRune(s[i:])
		reverse(s[i : i+size])
		i += size
	}
	reverse(s)
	return s
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// 正しくない
//func reverse(s []byte) []byte {
//	start := 0
//	end := len(s) - 1
//	for utf8.RuneCount(s[start:end+1]) > 1 {
//		_, fs := utf8.DecodeRune(s[start : end+1])
//		_, ls := utf8.DecodeLastRune(s[start : end+1])
//
//		for k := 0; k < fs || k < ls; k = k + 1 {
//			s[start+k], s[end-k] = s[end-k], s[start+k]
//			if k == ls-1 {
//				reverse2(s[start : start+k])
//			}
//			if k == fs-1 {
//				reverse2(s[end-k : end+1])
//			}
//		}
//		if fs > ls {
//			start += fs
//			end -= fs
//		} else {
//			start += ls
//			end -= ls
//		}
//	}
//	return s
//}
