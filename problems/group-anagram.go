package problems

func groupAnagrams(strs []string) [][]string {
	anagrams := map[[26]int][]string{}
	for _, str := range strs {
		count := [26]int{}
		for _, c := range str {
			count[c-'a']++
		}
		anagrams[count] = append(anagrams[count], str)
	}
	result := make([][]string, 0, len(anagrams))
	for _, anagram := range anagrams {
		result = append(result, anagram)
	}
	return result
}
