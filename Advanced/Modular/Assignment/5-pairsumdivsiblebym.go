package Assignment

func PairSumDivisiblebyM(A []int, B int) int {
	ans := 0
	freqMap := make(map[int]int)
	for i := range A {
		freqMap[A[i]%B] = freqMap[A[i]%B] + 1
	}

	ans += freqMap[0] * (freqMap[0] - 1) / 2
	l := 1
	r := B - 1

	for l < r {
		ans += freqMap[l] * freqMap[r]
		l++
		r--
	}

	if l == r {
		ans += freqMap[l] * (freqMap[l] - 1) / 2
	}

	return ans % (1e9 + 7)

}
