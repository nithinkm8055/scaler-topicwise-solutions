package Arrays

import (
	"fmt"
	"goplay/Advanced/Arrays/Arrays-1/Homework"
	"goplay/Advanced/Arrays/Arrays-2/Assignment"
)

func Arrays() {
	{
		fmt.Println(Homework.MaxSet([]int{25150, 1412, 82797, 48381, 7065, -47699, -25129, -65483, -64607, -45322, -55176, 27224, 80366, 60444, 70285, -93898, -41133, 96576, -58266, 21711, -20614, -95737, 20591, -48762, -76197, -38588, -54873, 37304, 61200, -68960, 93947}))
		fmt.Println(Homework.MaxSet([]int{0, 0, -1, 0}))
		fmt.Println(Homework.MaxSet([]int{-133242, -133243, 133242, -123}))

		fmt.Println(Homework.Flip("010"))

	}
	//
	{
		fmt.Println(Assignment.Submatrixsumqueries([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2}, []int{1, 2}, []int{2, 3}, []int{2, 3}))
	}
}
