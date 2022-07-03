package leetcode

// Write a program to solve a Sudoku puzzle by filling the empty cells.
// Empty cells are indicated by the character '.'.
// You may assume that there will be only one unique solution.
//
// Example 1:
//
// Input:
// [
//   ["5","3",".",".","7",".",".",".","."],
//   ["6",".",".","1","9","5",".",".","."],
//   [".","9","8",".",".",".",".","6","."],
//   ["8",".",".",".","6",".",".",".","3"],
//   ["4",".",".","8",".","3",".",".","1"],
//   ["7",".",".",".","2",".",".",".","6"],
//   [".","6",".",".",".",".","2","8","."],
//   [".",".",".","4","1","9",".",".","5"],
//   [".",".",".",".","8",".",".","7","9"]
// ]
// Output:
// [
//   ["5","3","4","6","7","8","9","1","2"],
//   ["6","7","2","1","9","5","3","4","8"],
//   ["1","9","8","3","4","2","5","6","7"],
//   ["8","5","9","7","6","1","4","2","3"],
//   ["4","2","6","8","5","3","7","9","1"],
//   ["7","1","3","9","2","4","8","5","6"],
//   ["9","6","1","5","3","7","2","8","4"],
//   ["2","8","7","4","1","9","6","3","5"],
//   ["3","4","5","2","8","6","1","7","9"]
// ]

type position struct {
	x int
	y int
}

func solveSudoku(board [][]byte) {
	pos, find := []position{}, false
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				pos = append(pos, position{x: i, y: j})
			}
		}
	}
	putSudoku(&board, pos, 0, &find)
}

func putSudoku(board *[][]byte, pos []position, index int, succ *bool) {
	if *succ == true {
		return
	}
	if index == len(pos) {
		*succ = true
		return
	}
	for i := 1; i < 10; i++ {
		if checkSudoku(board, pos[index], i) && !*succ {
			(*board)[pos[index].x][pos[index].y] = byte(i) + '0'
			putSudoku(board, pos, index+1, succ)
			if *succ == true {
				return
			}
			(*board)[pos[index].x][pos[index].y] = '.'
		}
	}
}

func checkSudoku(board *[][]byte, pos position, val int) bool {
	// 判断横行是否有重复数字
	for i := 0; i < len((*board)[0]); i++ {
		if (*board)[pos.x][i] != '.' && int((*board)[pos.x][i]-'0') == val {
			return false
		}
	}
	// 判断竖行是否有重复数字
	for i := 0; i < len((*board)); i++ {
		if (*board)[i][pos.y] != '.' && int((*board)[i][pos.y]-'0') == val {
			return false
		}
	}
	// 判断九宫格是否有重复数字
	posx, posy := pos.x-pos.x%3, pos.y-pos.y%3
	for i := posx; i < posx+3; i++ {
		for j := posy; j < posy+3; j++ {
			if (*board)[i][j] != '.' && int((*board)[i][j]-'0') == val {
				return false
			}
		}
	}
	return true
}

type Sudoku struct {
	row [9][10]bool
	col [9][10]bool
	box [9][10]bool
}

func (s *Sudoku) init() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			s.row[i][j] = false
			s.col[i][j] = false
			s.box[i][j] = false
		}
	}
}

func (s *Sudoku) print() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d ", s.row[i][j])}
		fmt.Println()
	}
}
func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	s := &Sudoku{}
	s.init(board)
	solveSudoku(s.board)
	s.print()
}