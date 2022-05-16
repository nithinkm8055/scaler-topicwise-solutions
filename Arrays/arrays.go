package Arrays

import (
	"fmt"
	Assignment3 "goplay/Arrays/CarryForward/Assignment"
	"goplay/Arrays/Intro/Assignment"
	"goplay/Arrays/Intro/Homework"
	Assignment2 "goplay/Arrays/PrefixSum/Assignment"
	Assignment4 "goplay/Arrays/SubArrays/Assignment"
	Homework2 "goplay/Arrays/SubArrays/Homework"
)

func Arrays() {

	{ // General Problems
		fmt.Println(Assignment.Rotate([]int{1, 2, 2}, 3))
		fmt.Println(Assignment.GoodPair([]int{1, 2, 3, 4}, 7))
		fmt.Println(Assignment.Maxmin([]int{1, 2, 3, 4}))
		fmt.Println(Homework.MaxElement([]int{2, 4, 3, 1, 5}, 3))
	}

	{ // Prefix Sum
		fmt.Println(Assignment2.PickfromBothSides([]int{5, -2, 3, 1, 2}, 3))
		fmt.Println(Assignment2.RangeSum([]int{1, 2, 3, 4, 5}, [][]int{{1, 4}, {2, 3}}))
	}

	{ // Carry Forward
		fmt.Println(Assignment3.SubsequenceAG("ABCGAG"))
		fmt.Println(Assignment3.ClosestMinMax([]int{2, 2, 3}))
	}

	{ // Subarrays
		fmt.Println(Assignment4.MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
		fmt.Println(Assignment4.SubarrayLeastAverage([]int{5, 15, 7, 6, 3, 4, 18, 14, 13, 7, 3, 7, 2, 18}, 6))
		fmt.Println(Homework2.AlternatingSubarrays([]int{0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 1}, 1))
	}

}
