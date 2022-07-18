package Assignment

// TLE
func solve(A []int) []int {
	result := make([]int, 0)
	for i := range A {
		result = append(result, numdivisors(A[i]))
	}
	return result
}

func numdivisors(a int) int {
	cnt := 0

	for i := 1; i*i <= a; i++ {

		if a%i == 0 {
			if a/i == i {
				cnt++
			} else {
				cnt += 2
			}
		}
	}

	return cnt

}
