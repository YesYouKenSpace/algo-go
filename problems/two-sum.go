package problems

func twoSum(nums []int, target int) []int {
	mapNums := map[int]int{}
	for i, num := range nums {
		if j, ok := mapNums[target-num]; ok {
			return []int{j, i}
		}
		mapNums[num] = i
	}
	return []int{}
}
