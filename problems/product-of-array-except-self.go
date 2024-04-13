package problems

func productExceptSelf(nums []int) []int {
    ans := make([]int, len(nums))

		indexWithZero := -1
		for i, num := range nums {
			if num == 0 {
				if indexWithZero == -1 {
					indexWithZero = i
				} else {
					return ans
				}
			}
		}

		if indexWithZero != -1 {
			ans[indexWithZero]=1
			for i, num:= range nums {
				if i == indexWithZero {
					continue
				}
				ans[indexWithZero] *= num
			}
			return ans
		}

		prefix :=1
		for i, num:= range nums {
			ans[i] = prefix
			prefix *= num
		}

		suffix :=1
		for i := len(nums)-1; i >=0; i-- {
			num := nums[i]
			ans[i] *= suffix
			suffix *= num
		}

		return ans
}