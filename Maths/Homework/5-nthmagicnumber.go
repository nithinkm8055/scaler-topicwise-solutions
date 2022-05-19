package Homework

func MagicNumber(A int) int {

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
		if count == A {
			break
		}
	}

	return magicNumber[A-1]

}
