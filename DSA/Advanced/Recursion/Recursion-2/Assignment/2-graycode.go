package Assignment

func grayCode(n int) []int {

	result := make([]int, 0)

	if n >= 1 {
		result = append(result, []int{0, 1}...)

	}

	for i := 2; i <= n; i++ {
		for j := len(result) - 1; j >= 0; j-- {
			result = append(result, result[j]+(1<<(i-1)))
		}
	}

	return result

}
