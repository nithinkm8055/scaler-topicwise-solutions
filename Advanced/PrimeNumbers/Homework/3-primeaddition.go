package Homework

func primeAddition(A int) int {
	if isPrime(A) {
		return 1
	}

	return 2
}

func isPrime(num int) bool {

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
