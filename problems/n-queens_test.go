package problems

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		n        int
		expected [][]string
	}{
		{
			n: 1,
			expected: [][]string{
				{"Q"},
			},
		},
		{
			n: 4,
			expected: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
		{
			n:        9,
			expected: [][]string{},
		},
	}

	for _, test := range tests {
		result := solveNQueens(test.n)
		if diff := cmp.Diff(test.expected, result); diff != "" {
			t.Errorf("solveNQueens(%d) mismatch (-expected +result):\n%s", test.n, diff)
		}
	}
}
