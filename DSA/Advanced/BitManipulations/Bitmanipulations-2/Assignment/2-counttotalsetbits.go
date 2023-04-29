package Assignment

func countTotalSetBits(A int) int {

	if A <= 1 {
		return A
	}
	x := setBitsPowerInTwoPower(A)
	btill2x := x * (1 << (x - 1))
	msbSet := A - (1 << x) + 1
	rest := A - (1 << x)

	ans := btill2x + msbSet + countTotalSetBits(rest)
	return ans % (1000000007)

}

func setBitsPowerInTwoPower(A int) int {
	x := 0

	for (1 << x) <= A {
		x++
	}

	return x - 1

}

func countTotalSetBits2(A int) int {
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
