package Homework

//optimized
func MagicNumber(A int) int {
	sum := 0
	number := 1
	for A > 0 {
		number *= 5
		if A%2 == 1 {
			sum += number
		}
		number /= 2
	}
	return sum
}

func MagicNumber2(A int) int {

	magicNumber := make([]int, 0)

	number := 1
	count := 0
	for count <= A {
		number = number * 5
		magicNumber = append(magicNumber, number)
		temp := len(magicNumber) - 1
		for i := 0; i < temp; i++ {
			count++
			magicNumber = append(magicNumber, magicNumber[i]+number)
		}
		count++
	}

	return magicNumber[A-1]
}
