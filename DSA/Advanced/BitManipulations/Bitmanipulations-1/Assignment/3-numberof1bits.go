package Assignment

func numSetBits(A int) int {

	cnt := 0
	for i := 0; i < 31; i++ {
		if A&(1<<i) != 0 {
			cnt++
		}
	}

	return cnt
}
