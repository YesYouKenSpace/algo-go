package problems

import (
	"regexp"
	"strings"
)

func ingest(data string) []string {
	splitted := strings.Split(data, ",")
	return splitted
}

func trim(s string) string {
	m := regexp.MustCompile(`(^")|("$)`)

	return strings.ReplaceAll(m.ReplaceAllString(s, ""), "\"\"", "\"")
	
}

func ingest2(data string) []string {
	ans := []string{}

	enabled := false
	from := 0
	for i, c := range data {
		if c == '"' {
			enabled = !enabled
		}
		if !enabled && c == ',' {
			ans = append(ans, trim(data[from:i]))
			from = i + 1
		}
	}
	ans = append(ans, trim(data[from:]))
	return ans
}
