package Homework

func divgame(A int, B int, C int) int {

	lcm := (B * C) / gcd(B, C)
	return A / lcm
}
