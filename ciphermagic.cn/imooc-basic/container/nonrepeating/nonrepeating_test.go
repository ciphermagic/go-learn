package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbb", 1},
		{"abcabcabcd", 4},

		// Chinese support
		{"一二三二一", 3},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("lengthOfNonRepeatingSubStr('%s'); got %d; expected %d", tt.s, actual, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "一二三二一"
	ans := 3
	actual := lengthOfNonRepeatingSubStr(s)
	for i := 0; i < b.N; i++ {
		if actual != ans {
			b.Errorf("lengthOfNonRepeatingSubStr('%s'); got %d; expected %d", s, actual, ans)
		}
	}
}
