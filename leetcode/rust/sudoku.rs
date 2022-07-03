fn sudoku(board: Vec<Vec<char>>) -> bool {
    let mut row = vec![vec![false; 9]; 9];
    let mut col = vec![vec![false; 9]; 9];
    let mut block = vec![vec![false; 9]; 9];
    for i in 0..9 {
        for j in 0..9 {
            if board[i][j] == '.' {
                continue;
            }
            let num = board[i][j] as u8 - '0' as u8;
            if row[i][num as usize - 1] || col[j][num as usize - 1] || block[i / 3 * 3 + j / 3][num as usize - 1] {
                return false;
            }
            row[i][num as usize - 1] = true;
            col[j][num as usize - 1] = true;
            block[i / 3 * 3 + j / 3][num as usize - 1] = true;
        }
    }
    true
}

fn output(board: &Vec<Vec<char>>) {
    for i in 0..9 {
        for j in 0..9 {
            print!("{}", board[i][j]);
        }
        println!();
    }
}

fn sudoku_solve_with_output(board: Vec<Vec<char>>) {
    for i in 0..9 {
        for j in 0..9 {
            if board[i][j] == '.' {
                for num in 1..=9 {
                    board[i][j] = (num as u8 + ('0' as u8)) as char;
                    if sudoku(board.clone()) {
                        output(&board);
                    }
                }
                board[i][j] = '.';
            }
        }
    }
}

fn main() {
    let board = vec![
        vec!['5', '3', '.', '.', '7', '.', '.', '.', '.'],
        vec!['6', '.', '.', '1', '9', '5', '.', '.', '.'],
        vec!['.', '9', '8', '.', '.', '.', '.', '6', '.'],
        vec!['8', '.', '.', '.', '6', '.', '.', '.', '3'],
        vec!['4', '.', '.', '8', '.', '3', '.', '.', '1'],
        vec!['7', '.', '.', '.', '2', '.', '.', '.', '6'],
        vec!['.', '6', '.', '.', '.', '.', '2', '8', '.'],
        vec!['.', '.', '.', '4', '1', '9', '.', '.', '5'],
        vec!['.', '.', '.', '.', '8', '.', '.', '7', '9'],
    ];
    println!("{}", sudoku_solve_with_output(board));
}