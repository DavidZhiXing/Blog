// you are given an n x n 2d matrix representing an image.
// rotate the image by 90 degrees (clockwise).
// you have to rotate the image in-place, which means you have to modify the input 2d matrix directly. DO
// NOT allocate another 2D matrix and do the rotation.
// example:
// Input:
// [
//   [1,2,3],
//   [4,5,6],
//   [7,8,9]
// ]
// Output:
// [
//   [7,4,1],
//   [8,5,2],
//   [9,6,3]

func rotate_clockwise(matrix [][]int) {
	length := len(matrix)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// rotate by vertical centerline 
	for i := 0; i < length; i++ {
		for j := 0; j < length/2; j++ {
			matrix[i][j], matrix[i][length-j-1] = matrix[i][length-j-1], matrix[i][j]
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate_clockwise(matrix)
	fmt.Println(matrix)
}