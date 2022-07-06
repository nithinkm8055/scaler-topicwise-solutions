package Assignment

func countTotalSetBits(A int) int {
	cnt := 0

	//TLE
	for i := 0; i < 32; i++ {
		for j := 1; j <= A; j++ {
			if j&(1<<i) != 0 {
				cnt++
			}
		}
	}

	return cnt
}

func binary(A int, cnt int) int {

	if A == 0 {
		return cnt
	}

	temp := A % 2
	if temp == 1 {
		cnt++
	}

	return binary(A/2, cnt)
}
