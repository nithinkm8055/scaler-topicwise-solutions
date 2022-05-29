package Maps

import (
	"fmt"
	"goplay/Maps/Day2/Assignment"
	"goplay/Maps/Day2/Homework"
)

func Maps() {
	{
		fmt.Println(Assignment.SubArraywithGivenSum([]int{1, 2, 3, 4, 5}, 5))
		fmt.Println(Assignment.TwoSum([]int{2, 7, 11, 15}, 9))
	}
	{
		fmt.Println(Homework.IsDictionary([]string{"fine", "nine", "nib"}, "qwertyuiopasdfghjklzxcvbnm"))
	}
}
