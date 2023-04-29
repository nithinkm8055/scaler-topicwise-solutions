package Assignment

func primesum(A int) []int {

	//result := make([]int, 0)

	for i := 2; i <= A/2; i++ {
		if checkPrime(i) && checkPrime(A-i) {
			return []int{i, A - i}
		}
	}

	return []int{}

}

func checkPrime(A int) bool {
	cnt := 0
	for i := 1; i*i <= A; i++ {
		if A%i == 0 {
			cnt++
			if i != A/i {
				cnt++
			}
			if cnt > 2 {
				return false
			}
		}
	}

	return true

}
