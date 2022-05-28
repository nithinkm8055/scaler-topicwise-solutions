package Contest1

import "fmt"

func Contest1() bool {
	var N, M, Q int
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &M)
	fmt.Scanf("%d", &Q)

	matrix(N, M, Q)

	return true
}

func matrix(N, M, Q int) {

	matrix := make([][]int, N)

	// forming matrix
	count := 1
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			matrix[i] = append(matrix[i], count)
			count++
		}
	}

	//fmt.Println(matrix)

	// read inputs

	for i := 0; i < Q; i++ {
		var query int
		fmt.Scanf("%d", &query)

		switch query {
		case 1:
			var c1 int
			var c2 int
			fmt.Scanf("%d", &c1)
			fmt.Scanf("%d", &c2)
			swapColumn(c1-1, c2-1, matrix)
		case 2:
			var r1 int
			var r2 int
			fmt.Scanf("%d", &r1)
			fmt.Scanf("%d", &r2)
			swapRow(r1-1, r2-1, matrix)
		case 3:
			indexes := make([]int, 4)
			for i := 0; i <= 3; i++ {
				fmt.Scanf("%d", &indexes[i])
			}
			BitwiseOR(indexes[0]-1, indexes[1]-1, indexes[2]-1, indexes[3]-1, matrix)
		case 4:
			indexes := make([]int, 4)
			for i := 0; i <= 3; i++ {
				fmt.Scanf("%d", &indexes[i])
			}
			BitwiseAND(indexes[0]-1, indexes[1]-1, indexes[2]-1, indexes[3]-1, matrix)
		}

	}
}

func swapColumn(col1, col2 int, matrix [][]int) {
	for row := 0; row < len(matrix); row++ {
		matrix[row][col1], matrix[row][col2] = matrix[row][col2], matrix[row][col1]
	}
}

func swapRow(row1, row2 int, matrix [][]int) {
	for col := 0; col < len(matrix[0]); col++ {
		matrix[row1][col], matrix[row2][col] = matrix[row2][col], matrix[row1][col]
	}
}

func BitwiseOR(x1, y1, x2, y2 int, matrix [][]int) {
	fmt.Println(matrix[x1][y1] | matrix[x2][y2])
}

func BitwiseAND(x1, y1, x2, y2 int, matrix [][]int) {
	fmt.Println(matrix[x1][y1] & matrix[x2][y2])
}
