package main

import (
	"encoding/hex"
	"fmt"
)

var (
	key   string = "ICE"
	plain string = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
)

func main() {
	fmt.Println(hex.EncodeToString(RepXOREncrypt([]byte(plain), []byte(key))))
}

func RepXOREncrypt(plain []byte, key []byte) []byte {
	plain_len := len(plain)
	key_len := len(key)
	result := make([]byte, plain_len)

	for i, j := 0, 0; i < plain_len; i, j = i+1, j+1 {
		if j == key_len {
			j = 0
		}
		result[i] = plain[i] ^ key[j]
	}

	return result
}
