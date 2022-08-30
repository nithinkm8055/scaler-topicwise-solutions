package Assignment

func addBinary(A string, B string) string {

	i := len(A) - 1
	j := len(B) - 1

	result := ""
	var carry uint8
	var a uint8

	var b uint8

	for i >= 0 || j >= 0 {

		if j < 0 {
			b = '0'
		} else {
			b = B[j]
		}

		if i < 0 {
			a = '0'
		} else {
			a = A[i]
		}

		sum := a + b + carry
		value := get(sum)

		result = value + result

		if sum > 97 {
			carry = 1
		} else {
			carry = 0
		}
		i--
		j--
	}

	if carry == 1 {
		value := "1"
		result = value + result
	}

	return result
}

func get(num uint8) string {

	if num == 97 || num > 98 {
		return "1"
	}

	return "0" // return if sum is 96 or 98
}
