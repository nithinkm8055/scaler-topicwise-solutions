package Assignment

func primeSubsequences(A []int) int {

	primeCnt := 0
	for i := range A {
		if isPrime(A[i]) {
			primeCnt++
		}
	}

	return pow(2, primeCnt, 1e9+7) - 1

}

func pow(A, B, C int) int {

	if A == 0 {
		return 0
	}

	if B == 0 {
		return 1
	}

	x := pow(A, B/2, C)

	if B%2 == 0 {

		return (x * x) % C
	}

	return ((A % C) * (x * x) % C) % C

}
