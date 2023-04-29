package Homework

import (
	"fmt"
)

func reversestr() {
	// YOUR CODE GOES HERE
	// Please take input and print output to standard input/output (stdin/stdout)
	// E.g. 'fmt.Scanf' for input & 'fmt.Printf' for output
	var input string
	fmt.Scanf("%s", &input)

	// old
	//fmt.Println(string(reverse([]byte(input), 0, len(input)-1)))
	reverse(input)
}

func reverse(s string) {

	if s == "" {
		return
	}

	fmt.Print(string(s[len(s)-1]))
	reverse(s[:len(s)-1])
}

func reverseOld(str []byte, i, j int) []byte {

	if i > j {
		return str
	}

	str[i], str[j] = str[j], str[i]

	return reverseOld(str, i+1, j-1)

}
