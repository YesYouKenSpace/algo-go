package problems

import (
	"fmt"
	"unicode"
)

// https://neetcode.io/problems/is-palindrome
// https://leetcode.com/problems/valid-palindrome/
func isPalindrome(s string) (bool, error) {
	left := 0
	right := len(s) - 1

	for left != right {
		if !isAlphanumeric(s[left]) {
			left += 1
			continue
		}
		if !isAlphanumeric(s[right]) {
			right -= 1
			continue
		}

		switch s[left] {
		case s[right]:
		case byte(unicode.ToLower(rune(s[right]))):
		case byte(unicode.ToUpper(rune(s[right]))):
		default:
			return false, fmt.Errorf("not matching: s[left]=%c s[right]=%c left=%d right=%d", s[left], s[right], left, right)
		}
		left += 1
		right -= 1
	}
	return true, nil
}

func isAlphanumeric(c byte) bool {
	switch {
	case c <= 'Z' && c >= 'A':
		return true
	case c <= 'z' && c >= 'a':
		return true
	case c <= '9' && c >= '0':
		return true
	default:
		return false
	}
}
