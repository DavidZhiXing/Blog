def validSudoku(board)
  # check rows
  board.each do |row|
    return false if !valid_row?(row)
  end
  true
end

def valid_row?(row)
    row.each do |num|
        return false if num == '.'
        return false if !num.match(/^[1-9]$/)
    end
    true
    end

def main
  board = [
    ["5","3",".",".","7",".",".",".","."],
    ["6",".",".","1","9","5",".",".","."],
    [".","9","8",".",".",".",".","6","."],
    ["8",".",".",".","6",".",".",".","3"],
    ["4",".",".","8",".","3",".",".","1"],
    ["7",".",".",".","2",".",".",".","6"],
    [".","6",".",".",".",".","2","8","."],
    [".",".",".","4","1","9",".",".","5"],
    [".",".",".",".","8",".",".","7","9"]
  ]
  puts validSudoku(board)
end

