package problems

import (
	"fmt"
)

// Back Tracking using while loop
// Time Complexity: O(n!)
// Space Complexity: O(n)
func solveNQueens(n int) [][]string {
	ans := [][]string{}
	queens := make([]int, n)

	str := make([]byte, n, n)
	for i := 0; i < n; i++ {
		str[i] = '.'
	}

	isColliding := func(q1, q2 int) bool {
		if queens[q1] == queens[q2] {
			return true
		}
		// check diagonal right +q2, +q2
		rowDiff := q2 - q1
		colDiff := queens[q2] - queens[q1]
		if colDiff == rowDiff || -colDiff == rowDiff {
			return true
		}
		return false
	}

	isViolating := func(queenNo int) bool {
		if queens[queenNo] >= n {
			return true
		}
		for i := 0; i < queenNo; i++ {
			if isColliding(i, queenNo) {
				return true
			}
		}
		return false
	}

	queenStrings := make([]string, n)
	addAns := func() {
		solution := make([]string, n)
		for i := 0; i < n; i++ {
			// fmt.Printf("Printing queen: %d pos: %d\n", i, queens[i])
			if queenStrings[queens[i]] == "" {
				queenStrings[queens[i]] = fmt.Sprintf("%s%s%s", str[:queens[i]], "Q", str[queens[i]+1:])
			}
			solution[i] = queenStrings[queens[i]]
		}
		// fmt.Printf("solution: %+v", solution)
		ans = append(ans, solution)
	}

	queenNo := 0

	backtrack := func() {
		queenNo -= 1
		queens[queenNo] += 1
	}

	for queens[0] < n {
		if queenNo == n {
			addAns()
			backtrack()
			continue
		}
		if !isViolating(queenNo) {
			queenNo += 1
			continue
		}
		queens[queenNo] += 1
		if queens[queenNo] >= n {
			queens[queenNo] = 0
			backtrack()
			continue
		}
	}
	return ans
}

// Brute force approach
// Time Complexity: O(n^n)
func solveNQueensUsingBruteForce(n int) [][]string {
	// Brute force
	// Each queen can only take up one row,
	// For each queen, try every column, compare with previous queens if they collide
	// try to solve for 1 first

	ans := [][]string{}
	queens := make([]int, n)

	str := make([]byte, n, n)
	for i := 0; i < n; i++ {
		str[i] = '.'
	}

	isColliding := func(q1, q2 int) bool {
		if queens[q1] == queens[q2] {
			return true
		}
		// check diagonal right +q2, +q2
		rowDiff := q2 - q1
		colDiff := queens[q2] - queens[q1]
		if colDiff == rowDiff || -colDiff == rowDiff {
			return true
		}
		return false
	}

	isViolating := func(queenNo int) bool {
		for i := 0; i < queenNo; i++ {
			if isColliding(i, queenNo) {
				return true
			}
		}
		return false
	}

	hasViolating := func() bool {
		for i := 0; i < n; i++ {
			if isViolating(i) {
				return true
			}
		}
		return false
	}

	next := func() bool {
		for i := 0; i < n; i++ {
			queens[i] += 1
			if queens[i] < n {
				return true
			}
			queens[i] = 0
		}
		return false
	}

	queenStrings := make([]string, n)
	addAns := func() {
		solution := make([]string, n)
		for i := 0; i < n; i++ {
			// var buffer bytes.Buffer
			if queenStrings[queens[i]] == "" {
				queenStrings[queens[i]] = fmt.Sprintf("%s%s%s", str[:queens[i]], "Q", str[queens[i]+1:])
			}
			solution[i] = queenStrings[queens[i]]
		}
		// fmt.Printf("solution: %+v", solution)
		ans = append(ans, solution)
	}

	hasNext := true

	for hasNext {
		if !hasViolating() {
			addAns()
		}
		hasNext = next()
	}

	return ans
}
