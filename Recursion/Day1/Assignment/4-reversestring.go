package Assignment

import "fmt"

func ReverseString() {
	var input string
	fmt.Scanf("%s", &input)
	fmt.Println(string(reverse([]byte(input), 0, len(input)-1)))
}

func reverse(bytes []byte, s, e int) []byte {

	if s >= e {
		return bytes
	}
	bytes[s], bytes[e] = bytes[e], bytes[s]
	return reverse(bytes, s+1, e-1)
}
