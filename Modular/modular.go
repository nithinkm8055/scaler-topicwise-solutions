package Modular

import (
	"fmt"
	"goplay/Modular/Assignment"
	"goplay/Modular/Homework"
)

func Modular() {

	{
		fmt.Println(Assignment.TitleToNumber("ZZZZZ"))
		fmt.Println(Assignment.Divisibilityby8("4758310052942341036336679074241759053154797537802772"))
		fmt.Println(Assignment.FindMod("43535321", 47))

		fmt.Println(Homework.RepeatedNumber([]int{3, 1, 2, 5, 3}))
		fmt.Println(Homework.Concatenate3Numbers(10, 20, 30))
		fmt.Println(Homework.RectangleOverlap(0, 0, 1, 1, 1, 1, 6, 6))
		fmt.Println(Homework.LCM(9, 6))
	}

}
