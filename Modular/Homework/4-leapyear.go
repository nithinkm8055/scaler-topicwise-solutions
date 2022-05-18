package Homework

func LeapYear(A int) int {
	if A%4 == 0 {
		if A%100 == 0 && A%400 == 0 {
			return 1
		} else if A%100 != 0 {
			return 1
		}
	}
	return 0
}
