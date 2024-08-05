package problems

import "testing"

func TestKthDistinct(t *testing.T) {
	tests := []struct {
		arr      []string
		k        int
		expected string
	}{
		{
			arr:      []string{"a", "b", "a", "c", "b", "d"},
			k:        1,
			expected: "c",
		},
		{
			arr:      []string{"a", "b", "a", "c", "b", "d"},
			k:        2,
			expected: "d",
		},
		{
			arr:      []string{"a", "b", "a", "c", "b", "d"},
			k:        3,
			expected: "",
		},
		{
			arr:      []string{"a", "b", "c", "d"},
			k:        1,
			expected: "a",
		},
		{
			arr:      []string{"a", "b", "c", "d"},
			k:        4,
			expected: "d",
		},
		{
			arr:      []string{"a", "a", "a"},
			k:        1,
			expected: "",
		},
	}

	for _, tc := range tests {
		actual := kthDistinct(tc.arr, tc.k)
		if actual != tc.expected {
			t.Errorf("expected: %v actual: %v, arr: %v, k: %d", tc.expected, actual, tc.arr, tc.k)
		}
	}
}
