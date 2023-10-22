package Homework

//Given a matrix of integers A of size N x M consisting of 0 and 1. A group of connected 1's forms an island. From a cell (i, j) such that A[i][j] = 1 you can visit any cell that shares a corner with (i, j) and value in that cell is 1.
//
//More formally, from any cell (i, j) if A[i][j] = 1 you can visit:
//
//(i-1, j) if (i-1, j) is inside the matrix and A[i-1][j] = 1.
//(i, j-1) if (i, j-1) is inside the matrix and A[i][j-1] = 1.
//(i+1, j) if (i+1, j) is inside the matrix and A[i+1][j] = 1.
//(i, j+1) if (i, j+1) is inside the matrix and A[i][j+1] = 1.
//(i-1, j-1) if (i-1, j-1) is inside the matrix and A[i-1][j-1] = 1.
//(i+1, j+1) if (i+1, j+1) is inside the matrix and A[i+1][j+1] = 1.
//(i-1, j+1) if (i-1, j+1) is inside the matrix and A[i-1][j+1] = 1.
//(i+1, j-1) if (i+1, j-1) is inside the matrix and A[i+1][j-1] = 1.
//Return the number of islands.
//
//NOTE: Rows are numbered from top to bottom and columns are numbered from left to right.

func solve(A [][]int) int {

	count := 0

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {

			if A[i][j] != 2 && A[i][j] != 0 {
				count++
				dfs(A, i, j)
			}
		}
	}
	return count
}

func dfs(A [][]int, i, j int) {

	A[i][j] = 2
	dx := []int{0, 0, -1, 1, 1, -1, 1, -1}
	dy := []int{1, -1, 0, 0, 1, 1, -1, -1}

	for d := 0; d < 8; d++ {
		x := i + dx[d]
		y := j + dy[d]

		if x >= 0 && x < len(A) && y >= 0 && y < len(A[0]) && A[x][y] == 1 {
			dfs(A, x, y)
		}
	}

}
