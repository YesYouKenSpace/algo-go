package problems

import "fmt"

func longestConsecutive(nums []int) int {
	store := map[int]struct{}{}
	for _, num := range nums {
		store[num] = struct{}{}
	}
	longest := 0
	for k := range store {
		if _, ok := store[k-1]; !ok {
			longest = max(longest, count(store, k))
		}
	}
	return longest
}

func count(store map[int]struct{}, num int) int {
	cnt := 0
	next := num
	for _, ok := (store[next+cnt]); ok; _, ok = (store[next+cnt]) {
		cnt++
	}
	return cnt
}

type TestCase struct {
	args     []int
	expected int
}

func main() {
	tests := []TestCase{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
	}

	for _, tt := range tests {
		actual := longestConsecutive(tt.args)
		if actual != tt.expected {
			fmt.Println(fmt.Errorf("longestConsecutive(%v): expected %d, actual %d\n", tt.args, tt.expected, actual))
		}
	}
}
