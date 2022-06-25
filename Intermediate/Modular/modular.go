package Modular

import (
	"fmt"
	Assignment2 "goplay/Intermediate/Modular/Assignment"
	Homework2 "goplay/Intermediate/Modular/Homework"
)

func Modular() {

	{
		fmt.Println(Assignment2.TitleToNumber("ZZZZZ"))
		fmt.Println(Assignment2.Divisibilityby8("4758310052942341036336679074241759053154797537802772"))
		fmt.Println(Assignment2.FindMod("43535321", 47))

		fmt.Println(Homework2.RepeatedNumber([]int{3, 1, 2, 5, 3}))
		fmt.Println(Homework2.Concatenate3Numbers(10, 20, 30))
		fmt.Println(Homework2.RectangleOverlap(0, 0, 1, 1, 1, 1, 6, 6))
		fmt.Println(Homework2.LCM(9, 6))
	}

}
