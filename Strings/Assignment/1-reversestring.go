package Assignment

import "strings"

func ReverseString(A string) string {

	reverse := reveseString(A)

	ans := ""
	index := 0
	for i := range reverse {
		if string(reverse[i]) == " " {
			ans = ans + reveseString(reverse[index:i]) + " "
			i++
			index = i
		} else if i == len(reverse)-1 {
			ans = ans + reveseString(reverse[index:])
		}
	}

	return strings.TrimSpace(ans)

}

func reveseString(A string) string {
	bytes := []byte(A)

	for i := 0; i < len(bytes)/2; i++ {
		bytes[i], bytes[len(bytes)-i-1] = bytes[len(bytes)-i-1], bytes[i]
	}
	return string(bytes)
}
