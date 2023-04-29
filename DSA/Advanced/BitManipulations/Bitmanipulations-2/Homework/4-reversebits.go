package Homework

func reverseBits(num uint32) uint32 {
	var reverse_num uint32
	for i := 0; i < 32; i++ {
		if (num & (1 << i)) != 0 {
			reverse_num |= 1 << ((32 - 1) - i)
		}
	}
	return reverse_num

}
