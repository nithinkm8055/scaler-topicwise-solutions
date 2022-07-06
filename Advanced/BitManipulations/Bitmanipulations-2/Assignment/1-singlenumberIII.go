package Assignment

func SingleNumberIII(A []int) []int {

	xor := 0
	for i := range A {
		xor = xor ^ A[i]
	}

	setIndex := -1
	for i := 0; i < 31; i++ {
		if xor&(1<<i) != 0 {
			setIndex = i
			break
		}
	}

	ans1 := 0
	ans2 := 0

	for i := range A {
		if A[i]&(1<<setIndex) != 0 {
			ans1 ^= A[i]
		} else {
			ans2 ^= A[i]
		}
	}

	if ans1 > ans2 {
		return []int{ans2, ans1}
	}

	return []int{ans1, ans2}
}
