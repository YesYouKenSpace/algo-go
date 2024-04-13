package problems

import "fmt"

func isValidSudoku(board [][]byte) bool {
	isValidRow := func(r int) bool {
		written := make(map[byte]struct{}, 9)
		for c := 0; c < 9; c++ {
			v := board[r][c]
			fmt.Printf("Row Check: r=%d c=%d, v=%s\n", r, c, string(v))
			if v == '.' {
				continue
			}
			if _, ok := written[v]; ok {
				fmt.Printf("Failed Row Check: r=%d c=%d, v=%s", r, c, string(v))
				return false
			}
			written[v] = struct{}{}
		}
		return true
	}

	isValidCol := func(c int) bool {
		written := make(map[byte]struct{}, 9)
		for r := 0; r < 9; r++ {

			v := board[r][c]
			fmt.Printf("Col Check: r=%d c=%d, v=%s\n", r, c, string(v))
			if v == '.' {
				continue
			}
			if _, ok := written[v]; ok {
				fmt.Println(fmt.Errorf("Failed Col Check: r=%d c=%d, v=%s", r, c, string(v)))
				return false
			}
			written[v] = struct{}{}
		}
		return true
	}

	isValidSquare := func(sqIndex int) bool {
		written := make(map[byte]struct{}, 9)
		for i := 0; i < 9; i++ {
			r := i%3 + 3*(sqIndex%3)
			c := i/3 + 3*(sqIndex/3)
			v := board[r][c]
			fmt.Printf("Sqr Check: r=%d c=%d, v=%s\n", r, c, string(v))
			if v == '.' {
				continue
			}
			if _, ok := written[v]; ok {
				fmt.Printf("Failed Sqr Check: r=%d c=%d, v=%s", r, c, string(v))
				return false
			}
			written[v] = struct{}{}
		}
		return true
	}

	for i := 0; i < 9; i++ {
		if !isValidCol(i) {
			return false
		}
		if !isValidRow(i) {
			return false
		}
		if !isValidSquare(i) {
			return false
		}
	}
	return true
}
