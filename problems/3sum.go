package main

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	var result [][]int

	for a := 0; a < n-2; a++ {
		b := a + 1
		c := n - 1
		for b < c {
			sum := nums[a] + nums[b] + nums[c]
			if sum == 0 {
				result = append(result, []int{nums[a], nums[b], nums[c]})
			}
			if sum < 0 {
				b++
			} else {
				c--
				for b < c && nums[c] == nums[c+1] {
					c--
				}
			}
		}
		for a < n-2 && nums[a] == nums[a+1] {
			a++
		}
	}
	return result
}
