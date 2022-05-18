package Homework

// this biggy was unfortunately my initial approach
func RectangleOverlap(A int, B int, C int, D int, E int, F int, G int, H int) int {
	//AC & EG should have a number in common
	//BD & FH should have a number in common

	xAxisA := updateAxisElements(A, C)
	yAxisA := updateAxisElements(B, D)
	xAxisE := updateAxisElements(E, G)
	yAxisE := updateAxisElements(F, H)

	if checkCommon(xAxisA, xAxisE) && checkCommon(yAxisA, yAxisE) {
		return 1
	}

	return 0
}

func checkCommon(A []int, B []int) bool {
	for i := range A {
		for j := 0; j < len(B); j++ {
			if A[i] == B[j] {
				return true
			}
		}
	}
	return false
}

func updateAxisElements(A int, B int) []int {
	min := A
	max := A
	if A > max {
		min = B
		max = A
	} else {
		min = A
		max = B
	}
	temp := make([]int, 0)
	for i := min; i < max; i++ {
		temp = append(temp, i)
	}

	return temp
}
