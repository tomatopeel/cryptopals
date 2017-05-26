package bitutils

import (
	"errors"
)

func Hamming(a, b []byte) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("passed []bytes of unequal length")
	}
	c := 0
	for i := 0; i < len(a); i++ {
		c += CountBits(int(a[i] ^ b[i]))
	}
	return c, nil
}

func CountBitsNaive(x int) int {
	c := 0
	for x > 0 {
		c += x & 1
		x >>= 1
	}
	return c
}

func CountBitsKernighan(x int) int {
	c := 0
	for x > 0 {
		x &= x - 1
		c++
	}
	return c
}

func CountBits(v int) int {
	c := 0
	c = v - ((v >> 1) & 0x55555555)
	c = ((c >> 2) & 0x33333333) + (c & 0x33333333)
	c = ((c >> 4) + c) & 0x0F0F0F0F
	c = ((c >> 8) + c) & 0x00FF00FF
	c = ((c >> 16) + c) & 0x0000FFFF
	return c
}
