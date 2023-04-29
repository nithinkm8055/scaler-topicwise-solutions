package Homework

import "math"

func NumberOfOpenDoors(A int) int {
	// all the door numbers that are perfect squares remain at the end
	return int(math.Sqrt(float64(A)))
}
