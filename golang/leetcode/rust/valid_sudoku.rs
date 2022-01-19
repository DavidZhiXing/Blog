fn validSudoku(board: Vec<Vec<char>>) -> bool {
    let mut row = vec![vec![false; 9]; 9];
    let mut col = vec![vec![false; 9]; 9];
    let mut block = vec![vec![false; 9]; 9];
    for i in 0..9 {
        for j in 0..9 {
            let c = board[i][j];
            if c == '.' {
                continue;
            }
            let c = c as usize - '1' as usize;
            if row[i][c] || col[j][c] || block[i / 3 * 3 + j / 3][c] {
                return false;
            }
            row[i][c] = true;
            col[j][c] = true;
            block[i / 3 * 3 + j / 3][c] = true;
        }
    }
    true
}

fn main() {
    test_validSudoku();

#[test]
fn test_validSudoku() {
    assert_eq!(validSudoku(vec![
        vec!['5', '3', '.', '.', '7', '.', '.', '.', '.'],
        vec!['6', '.', '.', '1', '9', '5', '.', '.', '.'],
        vec!['.', '9', '8', '.', '.', '.', '.', '6', '.'],
        vec!['8', '.', '.', '.', '6', '.', '.', '.', '3'],
        vec!['4', '.', '.', '8', '.', '3', '.', '.', '1'],
        vec!['7', '.', '.', '.', '2', '.', '.', '.', '6'],
        vec!['.', '6', '.', '.', '.', '.', '2', '8', '.'],
        vec!['.', '.', '.', '4', '1', '9', '.', '.', '5'],
        vec!['.', '.', '.', '.', '8', '.', '.', '7', '9'],
    ]), true);
    assert_eq!(validSudoku(vec![
        vec!['8', '3', '.', '.', '7', '.', '.', '.', '.'],
        vec!['6', '.', '.', '1', '9', '5', '.', '.', '.'],
        vec!['.', '9', '8', '.', '.', '.', '.', '6', '.'],
        vec!['8', '.', '.', '.', '6', '.', '.', '.', '3'],
        vec!['4', '.', '.', '8', '.', '3', '.', '.', '1'],
        vec!['7', '.', '.', '.', '2', '.', '.', '.', '6'],
        vec!['.', '6', '.', '.', '.', '.', '2', '8', '.'],
        vec!['.', '.', '.', '4', '1', '9', '.', '.', '5'],
        vec!['.', '.', '.', '.', '8', '.', '.', '7', '9'],
    ]), false);