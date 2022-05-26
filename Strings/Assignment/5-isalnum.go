package Assignment

func Isalnum(A []byte) int {
	for i := range A {
		if (A[i] >= 65 && A[i] <= 90) || (A[i] >= 97 && A[i] <= 122) || (A[i] >= '0' && A[i] <= '9') {
			continue
		}
		return 0
	}

	return 1

}
