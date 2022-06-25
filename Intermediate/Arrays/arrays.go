package Arrays

import (
	"fmt"
	Assignment7 "goplay/Intermediate/Arrays/CarryForward/Assignment"
	Assignment5 "goplay/Intermediate/Arrays/Intro/Assignment"
	"goplay/Intermediate/Arrays/Intro/Homework"
	Assignment3 "goplay/Intermediate/Arrays/PrefixSum/Assignment"
	Assignment6 "goplay/Intermediate/Arrays/SubArrays/Assignment"
	Homework2 "goplay/Intermediate/Arrays/SubArrays/Homework"
)

func Arrays() {

	{ // General Problems
		fmt.Println(Assignment5.Rotate([]int{1, 2, 2}, 3))
		fmt.Println(Assignment5.GoodPair([]int{1, 2, 3, 4}, 7))
		fmt.Println(Assignment5.Maxmin([]int{1, 2, 3, 4}))
		fmt.Println(Homework.MaxElement([]int{2, 4, 3, 1, 5}, 3))
	}

	{ // Prefix Sum
		fmt.Println(Assignment3.PickfromBothSides([]int{5, -2, 3, 1, 2}, 3))
		fmt.Println(Assignment3.RangeSum([]int{1, 2, 3, 4, 5}, [][]int{{1, 4}, {2, 3}}))
	}

	{ // Carry Forward
		fmt.Println(Assignment7.SubsequenceAG("ABCGAG"))
		fmt.Println(Assignment7.ClosestMinMax([]int{2, 2, 3}))
	}

	{ // Subarrays
		fmt.Println(Assignment6.MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
		fmt.Println(Assignment6.SubarrayLeastAverage([]int{5, 15, 7, 6, 3, 4, 18, 14, 13, 7, 3, 7, 2, 18}, 6))
		fmt.Println(Homework2.AlternatingSubarrays([]int{0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 1}, 1))
	}

}
