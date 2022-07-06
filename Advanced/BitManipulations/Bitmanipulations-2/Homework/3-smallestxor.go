package Homework

func solve(A int, B int) int {
	ans := 0
	setbit := 0
	flag := false

	// set the MSB in the ans based on A, this helps to minimise xor
	for j := 32; j >= 0; j-- {
		if A&(1<<j) != 0 {
			ans |= (1 << j)
			setbit++
		}
		if setbit == B {
			flag = true
			break
		}
	}

	// fmt.Println("Ans :", ans)
	// set the lsb in ans which are not set
	if !flag {
		for j := 0; j < 32; j++ {

			if ans&(1<<j) == 0 {
				ans |= (1 << j)
				setbit++
			}

			if setbit == B {
				break
			}

		}
	}
	return ans
}
