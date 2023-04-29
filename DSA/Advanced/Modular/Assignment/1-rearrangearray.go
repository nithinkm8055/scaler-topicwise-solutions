package Assignment

func arrange(A []int) {

	N := len(A)
	for i := range A {
		A[i] = A[i] * N
	}

	for i := range A {
		idx := A[i] / N
		newValue := A[idx] / N // division to cancel out the multiplication done in first loop
		A[i] += newValue
	}

	for i := range A {
		A[i] %= N
	}

}
