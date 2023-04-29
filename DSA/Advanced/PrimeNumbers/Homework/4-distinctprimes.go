package Homework

func distinctPrimes(A []int) int {

	countMap := make(map[int]int)
	for i := range A {
		countPrimedivisors(A[i], countMap)
	}

	return len(countMap)

}

func countPrimedivisors(A int, countMap map[int]int) {
	for i := 1; i*i <= A; i++ {
		if A%i == 0 {
			if isPrime(i) {
				countMap[i] = countMap[i] + 1
			}

			if i != A/i && isPrime(A/i) {
				countMap[A/i] = countMap[A/i] + 1
			}
		}
	}

}
