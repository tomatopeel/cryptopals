package datautils

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

func RepeatingKeyXOR(data []byte, key []byte) []byte {
	data_len := len(data)
	key_len := len(key)
	result := make([]byte, data_len)

	for i, j := 0, 0; i < data_len; i, j = i+1, j+1 {
		if j == key_len {
			j = 0
		}
		result[i] = data[i] ^ key[j]
	}

	return result
}

// Split data into k len []byte's and return the [][]byte
func Blocks(k int, data []byte) (blocks [][]byte) {
	for b, rem := Block(k, data); b != nil; b, rem = Block(k, rem) {
		blocks = append(blocks, b)
	}
	return
}

// Return k length []byte and the remaining []byte
func Block(k int, data []byte) ([]byte, []byte) {
	if len(data) < k {
		return nil, nil
	}
	return data[:k], data[k:]
}
