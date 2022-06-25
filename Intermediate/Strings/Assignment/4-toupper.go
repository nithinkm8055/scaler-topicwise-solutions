package Assignment

func To_upper(A []byte) []byte {

	for i := range A {
		if A[i] >= 97 && A[i] <= 122 {
			A[i] = A[i] - 32
		}
	}

	return A
}
