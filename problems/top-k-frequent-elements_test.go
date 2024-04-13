package problems

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTopKFrequent(t *testing.T) {
	testCases := []struct{
		nums []int
		k int
		expected []int
	}{
	// 	{
	// 	nums: []int{1,1,1,2,2,3},
	// 	k: 2,
	// 	expected: []int{1, 2},
	// },
	// {
	// 	nums: []int{5,3,1,1,1,3,73,1},
	// 	k: 2,
	// 	expected: []int{1,3},
	// },
	{
		nums: []int{4,1,-1,2,-1,2,3},
		k: 2,
		expected: []int{-1,2},
	},
}

	for _,tt := range testCases {
		res := topKFrequent(tt.nums, tt.k)
		if diff := cmp.Diff(res, tt.expected); diff != "" {
			t.Errorf("Got: %v\nExpected: %v", res, tt.expected)
		}
	}
}