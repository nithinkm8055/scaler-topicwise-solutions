package Homework

func PairswithGivenXOR(A []int, B int) int {
	// if a ^ b = c , then a ^ c = b, search for b in xorMap
	// already given slice is distince

	xorMap := make(map[int]int)
	count := 0
	for i := range A {
		C := A[i] ^ B
		if _, ok := xorMap[A[i]]; ok {
			count++
		} else {
			xorMap[C] = 0
		}

	}
	return count
}
