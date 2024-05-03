package main

import "fmt"

var TakeStep = "X"
var SkipStep = "-"

func getCombi(h int, res [][]string) [][]string {
	if h == 0 {
		return res
	}
	res2 := [][]string{}
	for _, arr := range res {
		arr1 := make([]string, len(arr))
		copy(arr1, arr)
		res2 = append(res2, append(arr1, TakeStep))
		if len(arr) == 0 || arr[len(arr)-1] == TakeStep && h > 1 {
			arr2 := make([]string, len(arr))
			copy(arr2, arr)
			res2 = append(res2, append(arr2, SkipStep))
		}
	}
	return getCombi(h-1, res2)
}

func printRes(ans [][]string) {
	for _, a := range ans {
		fmt.Println(a)
	}
}

func main() {
	ans := getCombi(10, [][]string{{}})
	for _, a := range ans {
		fmt.Println(a)
	}
	fmt.Println("ans: ", len(ans))

	for i := 1; i <= 10; i++ {
		fmt.Println(i, ":", getNum(i))
	}
}

// 2: - x
// 2: x, x

// 3: x, x, x
// 3: -, x, x
// 3: x, -, x

// 4: x, x, x, x
// 4: -, x, x, x
// 4: -, x, -, x
// 4: x, -,x, x
// 4: x, x, -,x

// 5: x, x, x, x
// 5: -, x, x, x
// 5: -, x, -, x
// 5: x, -,x, x
// 5: x, x, -,x

func getNum(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 2 {
		return n
	}
	return getNum(n-1) + getNum(n-2)
}
