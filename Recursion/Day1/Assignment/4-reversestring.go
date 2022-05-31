package Assignment

func reverse(A string, s, e int) string {
	if s >= e {
		return A
	}

	bytes := []byte(A)
	bytes[s], bytes[e] = bytes[e], bytes[s]
	
	return reverse(string(bytes), s+1, e-1)
}
