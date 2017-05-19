package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

var (
	hex_secret string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	alphabet   string = "abcdefghijklmnopqrstuvwxyz"
)

func main() {
	secret, err := hex.DecodeString(hex_secret)
	if err != nil {
		log.Fatal(err)
	}
	for _, char := range alphabet {
		result := isIt(char, secret)
		if result != "" {
			fmt.Println(result)
			break
		}
	}
}

func isIt(char int32, secret []byte) string {
	tester := make([]byte, len(secret))
	for i := 0; i < len(secret); i++ {
		tester[i] = byte(char)
	}
	for i := range secret {
		tester[i] ^= secret[i]
	}
	fmt.Println(string(tester))
	return ""
}

func counter_letters(word string) map[rune]int {
	return nil
}
