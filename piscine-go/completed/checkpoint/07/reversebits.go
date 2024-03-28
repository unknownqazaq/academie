package _7

func ReverseBits(oct byte) byte {
	result := byte(0)
	intSize := 7 // Assuming a byte has 8 bits
	for oct != 0 {
		result += (oct & 1) << intSize
		oct = oct >> 1
		intSize -= 1
	}
	return result
}
