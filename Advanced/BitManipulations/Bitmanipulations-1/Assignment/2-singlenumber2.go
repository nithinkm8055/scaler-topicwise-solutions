package Assignment

func SingleNumber2(A []int) int {
	ans := 0
	for i := 0; i < 31; i++ {

		cnt := 0

		for j := 0; j < len(A); j++ {
			if A[j]&(1<<i) != 0 {
				cnt++
			}
		}

		if cnt%3 != 0 {
			ans = ans | (1 << i)
		}

	}

	return ans
}
