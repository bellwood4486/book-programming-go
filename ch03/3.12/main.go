package anagram

import (
	"sort"
	"strings"
)

func isAnagram(s1, s2 string) bool {
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)
	for _, r := range s1 {
		m1[r]++
	}
	for _, r := range s2 {
		m2[r]++
	}

	if len(m1) != len(m2) {
		return false
	}
	for r, n := range m1 {
		if m2[r] != n {
			return false
		}
	}
	return true
}

func isAnagram2(s1, s2 string) bool {
	s1a := strings.Split(s1, "")
	s2a := strings.Split(s2, "")
	sort.Strings(s1a)
	sort.Strings(s2a)
	return strings.Join(s1a, "") == strings.Join(s2a, "")
}
