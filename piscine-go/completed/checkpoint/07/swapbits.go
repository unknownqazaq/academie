package _7

func SwapBits(octet byte) byte {
	// Swap the bits using bitwise operations
	return (octet >> 4) | (octet << 4)
}
