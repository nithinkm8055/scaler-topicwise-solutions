package Day1

import (
	"fmt"
	"goplay/Maps/Day1/Assignment"
	"goplay/Maps/Day1/Homework"
)

func Maps() {
	{
		fmt.Println(Assignment.Lszero([]int{22, -7, 15, -21, -20, -8, 16, -20, 5, 2, -15, -24, 19, 25, 3, 23, 18, 0, 16, 26, 13, 4, -20, -13, -25, -2}))
	}
	{
		fmt.Println(Homework.CheckPalindrome("abcde"))
		fmt.Println(Homework.Colorful(236))
	}
}
