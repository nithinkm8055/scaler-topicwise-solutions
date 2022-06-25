package Assignment

func Abmodulo(A int, B int) int {
	c := A - B
	if c < 0 {
		c = c * -1
	}
	return c
}
