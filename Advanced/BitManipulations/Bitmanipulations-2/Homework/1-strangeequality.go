package Homework

func StrangeEquality(A int) int {

	X := toggle(A)
	Y := closestPowerOf2GreaterThanA(A)

	return X ^ (1 << Y)

}

func closestPowerOf2GreaterThanA(A int) int {

	ans := 0

	for (1 << ans) <= A {
		ans++
	}

	return ans

}

func toggle(A int) int {

	msb := findMSB(A)

	for i := 0; i <= msb; i++ {
		A = A ^ (1 << i)
	}

	return A

}

func findMSB(A int) int {
	ans := 0
	for i := 0; i < 31; i++ {
		if A&(1<<i) != 0 {
			ans = i
		}
	}

	return ans

}
