package Maps

import (
	"fmt"

	Assignment2 "github.com/nithinkm8055/scaler-topicwise-solutions/Intermediate/Maps/Day2/Assignment"
	"github.com/nithinkm8055/scaler-topicwise-solutions/Intermediate/Maps/Day2/Homework"
)

func Maps() {
	{
		fmt.Println(Assignment2.SubArraywithGivenSum([]int{1, 2, 3, 4, 5}, 5))
		fmt.Println(Assignment2.TwoSum([]int{2, 7, 11, 15}, 9))
	}
	{
		fmt.Println(Homework.IsDictionary([]string{"fine", "nine", "nib"}, "qwertyuiopasdfghjklzxcvbnm"))
	}
}
