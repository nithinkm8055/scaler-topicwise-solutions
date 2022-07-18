package Assignment

func luckyNumbers(A int) int {

	cnt := 0
	for i := 2; i <= A; i++ {
		div := 0
		flag := true
		for j := 1; j*j <= i; j++ {
			if i%j == 0 {
				if isPrime(j) {
					div++
				}

				if j != i/j && isPrime(i/j) {
					div++
				}
			}
		}
		if div == 2 && flag {
			cnt++
		}

	}

	return cnt
}

func isPrime(A int) bool {
	cnt := 0
	for i := 1; i*i <= A; i++ {
		if A%i == 0 {
			cnt++
			if i != A/i {
				cnt++
			}
		}
	}
	if cnt == 2 {
		return true
	}
	return false
}
