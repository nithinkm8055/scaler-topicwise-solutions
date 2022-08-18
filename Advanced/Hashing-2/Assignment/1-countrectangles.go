package Assignment

func CountRectangles(A []int, B []int) int {

	type pair struct {
		x int
		y int
	}
	axisMap := make(map[pair]int)

	for i := range A {
		axisMap[pair{A[i], B[i]}]++
	}
	ans := 0

	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if A[i] != A[j] && B[i] != B[j] {
				if axisMap[pair{A[i], B[j]}] == 1 && axisMap[pair{A[j], B[i]}] == 1 {
					ans++
				}
			}

		}
	}
	return ans / 2
}
