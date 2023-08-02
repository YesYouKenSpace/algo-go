package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	char_map := map[rune]int16{}
	for _, c := range s {
		char_map[c]++
	}
	for _, c := range t {
		if char_map[c] <= 0 {
			return false
		}
		char_map[c]--
	}
	return true
}
