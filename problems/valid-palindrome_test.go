package problems

import "testing"

// https://neetcode.io/problems/is-palindrome
// https://leetcode.com/problems/valid-palindrome/
func TestIsPalindrome(t *testing.T) {

	tests := []struct {
		arg      string
		expected bool
	}{
		{
			arg:      "Was it a car or a cat I saw?",
			expected: true,
		},
		{
			arg:      "tab a cat",
			expected: false,
		},
		{
			arg:      "aa",
			expected: true,
		},
	}

	for _, tc := range tests {
		actual, err := IsPalindrome(tc.arg)
		if actual != tc.expected {
			t.Errorf("expected: %v actual: %v, arg: %s, err: %+v", tc.expected, actual, tc.arg, err)
		}
	}
}
