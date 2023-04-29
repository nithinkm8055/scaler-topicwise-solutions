package Assignment

func ncrmodm(A, B, K int) int {

	C := make([][]int, A+1)
	for i := 0; i <= A; i++ {
		C[i] = make([]int, B+1)
		for j := 0; j <= B; j++ {
			if j == 0 {
				C[i][j] = 1
			} else if j > i {
				C[i][j] = 0
			} else if i == j {
				C[i][j] = 1
			} else {
				C[i][j] = ((C[i-1][j-1] % K) + (C[i-1][j] % K)) % K
			}
		}
	}
	return C[A][B] % K
}

// Recursion
// func solve(n int , r int , m int )  (int) {
//     if r == n || r == 0 {
// 		return 1
// 	}

// 	if r == 1 {
// 		return n
// 	}

// 	return ((solve(n-1, r-1, m) % m) + (solve(n-1, r, m) % m)) % m
// }
