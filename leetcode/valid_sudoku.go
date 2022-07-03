package leetcode

// determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.
// The Sudoku board could be partially filled, where empty cells are filled with the character '.'.
// A partially filled sudoku which is valid.
// Note:
// A valid Sudoku board (partially filled) is not necessarily solvable. Only the filled cells need to be validated.
// The given board contain only digits 1-9 and the character '.'.
// The given board size is always 9x9.

type Sudoku struct {
	row [9][10]bool
	col [9][10]bool
	box [9][10]bool
}

func (s *Sudoku) check(i, j, num int) bool {
	if s.row[i][num] || s.col[j][num] || s.box[i/3*3+j/3][num] {
		return false
	}
	s.row[i][num] = true
	s.col[j][num] = true
	s.box[i/3*3+j/3][num] = true
	return true
}

// test
func (s *Sudoku) reset(i, j, num int) {
	s.row[i][num] = false
	s.col[j][num] = false
	s.box[i/3*3+j/3][num] = false
}

func (s *Sudoku) isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := int(board[i][j] - '0')
			if !s.check(i, j, num) {
				return false
			}
		}
	}
	return true
}

