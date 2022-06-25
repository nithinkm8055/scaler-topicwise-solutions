package Assignment

func To_lower(A []byte) []byte {
	for i := range A {
		if A[i] >= 65 && A[i] <= 90 {
			A[i] = A[i] + 32
		}
	}

	return A
}
