package Assignment

func permute(A []int) [][]int {
	result := make([][]int, 0)
	permutations(0, A, &result)
	return result
}

func permutations(i int, arr []int, result *[][]int) {

	if i == len(arr) {
		b := make([]int, 0)
		b = append(b, arr...)
		*result = append(*result, b)
		return
	}

	for j := i; j < len(arr); j++ {
		arr[i], arr[j] = arr[j], arr[i]
		permutations(i+1, arr, result)
		arr[i], arr[j] = arr[j], arr[i]
	}

}
