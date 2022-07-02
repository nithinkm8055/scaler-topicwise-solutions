package Assignment

func maxArr(A []int) int {

	mx1 := -1000000001
	mx2 := -1000000001
	mn1 := 1000000001
	mn2 := 1000000001

	for i := range A {

		mx1 = max(mx1, A[i]+i)
		mn1 = min(mn1, A[i]+i)

		mx2 = max(mx2, A[i]-i)
		mn2 = min(mn2, A[i]-i)
	}

	return max(mx1-mn1, mx2-mn2)

}
