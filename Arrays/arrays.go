package Arrays

import (
	"fmt"
	Assignment3 "goplay/Arrays/CarryForward/Assignment"
	"goplay/Arrays/Intro/Assignment"
	"goplay/Arrays/Intro/Homework"
	Assignment2 "goplay/Arrays/PrefixSum/Assignment"
)

func Arrays() {

	fmt.Println("######################### G E N E R A L ##############################")
	{ // General Problems
		fmt.Println(Assignment.Rotate([]int{1, 2, 2}, 3))
		fmt.Println(Assignment.GoodPair([]int{1, 2, 3, 4}, 7))
		fmt.Println(Assignment.Maxmin([]int{1, 2, 3, 4}))
		fmt.Println(Homework.MaxElement([]int{2, 4, 3, 1, 5}, 3))
	}
	fmt.Println("######################### P R E F I X S U M ##############################")
	{ //Prefix Sum
		fmt.Println(Assignment2.PickfromBothSides([]int{5, -2, 3, 1, 2}, 3))
		fmt.Println(Assignment2.RangeSum([]int{1, 2, 3, 4, 5}, [][]int{{1, 4}, {2, 3}}))
	}
	fmt.Println("######################### C A R R Y F O R W A R D ##############################")
	{
		fmt.Println(Assignment3.SubsequenceAG("ABCGAG"))
		fmt.Println(Assignment3.ClosestMinMax([]int{2, 2, 3}))
	}

}
