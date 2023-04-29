package Discussion

import "fmt"

func StarPattern() {

	var input int
	fmt.Scanf("%d", &input)

	for i := 0; i < input; i++ {
		for j := 0; j < 2*input; j++ {
			if j < input-i || (j >= input+i) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	for i := 0; i < input; i++ {
		for j := 0; j < 2*input; j++ {
			if j <= i || j >= 2*input-i-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}
