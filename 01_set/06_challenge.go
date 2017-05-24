package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	remote_file string = "http://cryptopals.com/static/challenge-data/6.txt"
	local_file  string = "secrets_01_06.txt"
	ham_a       string = "this is a test"
	ham_b       string = "wokka wokka!!!"
)

func main() {
	result, err := ComputeHamming([]byte(ham_a), []byte(ham_b))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func ComputeHamming(a, b []byte) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("ComputeHamming: was passed strings of differing lengths")
	}
	count := 0
	for i, _ := range a {
		a[i] ^= b[i]
		count += CountBits(int(a[i]))
	}
	return count, nil
}

func CountBitsNaive(x int) int {
	count := 0
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	if count == 0 {
		log.Fatal("Hm")
	}
	return count
}

func CountBitsKernighan(x int) int {
	count := 0
	for x > 0 {
		x &= x - 1
		count++
	}
	return count
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
