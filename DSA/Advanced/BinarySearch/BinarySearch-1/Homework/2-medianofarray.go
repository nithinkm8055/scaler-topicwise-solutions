package Homework

// TLE submission
// TC(O(max(n,m)), SC(O(n+m))
func findMedianSortedArrays(A []int, B []int) float64 {
	i := 0
	j := 0
	// merge arrays
	C := make([]int, len(A)+len(B))
	k := 0
	for i < len(A) && j < len(B) {

		if A[i] <= B[j] {
			C[k] = A[i]
			k++
			i++
		} else {
			C[k] = B[j]
			k++
			j++
		}
	}

	for i < len(A) {
		C[k] = A[i]
		k++
		i++
	}

	for j < len(B) {
		C[k] = B[j]
		k++
		j++
	}

	if len(C)%2 != 0 {
		idx := len(C) / 2
		return float64(C[idx])
	}

	idx1 := len(C) / 2
	idx2 := (len(C) / 2) - 1

	return float64(C[idx1]+C[idx2]) / 2
}
