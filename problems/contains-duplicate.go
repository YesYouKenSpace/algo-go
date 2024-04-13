package problems

func containsDuplicate(nums []int) bool {
	nums_map := map[int]byte{}
	for _, num := range nums {
		if _, ok := nums_map[num]; ok {
			return true
		} else {
			nums_map[num] = 1
		}
	}
	return false
}
