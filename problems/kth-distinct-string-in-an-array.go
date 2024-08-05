package problems

func kthDistinct(arr []string, k int) string {
	appeared := map[string]struct{}{}
	hasDuplicates := map[string]struct{}{}
	for _, a := range arr {
		if _, ok := appeared[a]; ok {
			hasDuplicates[a] = struct{}{}
		} else {
			appeared[a] = struct{}{}
		}
	}

	count := 0
	for _, a := range arr {
		if _, ok := hasDuplicates[a]; !ok {
			count+=1
			if count == k {
				return a
			}
		}
	}
	return ""
}
