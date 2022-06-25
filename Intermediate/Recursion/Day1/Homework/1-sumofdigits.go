package Homework

func Sum(A int) int {

	return sumofDigits(A, 0)

}

func sumofDigits(A int, sum int) int {
	if A <= 0 {
		return sum
	}
	temp := A % 10
	A = A / 10

	sum += temp

	return sumofDigits(A, sum)
}
