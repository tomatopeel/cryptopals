package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

var (
	hex_secret string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	alphabet   string = "abcdefghijklmnopqrstuvwxyz"
	vowels     string = "aeiouAEIOU"
	consonants string = "bcdfghjklmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ"
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
	fmt.Printf("%d\t%s\n", score(tester), string(tester))
	return ""
}

func score(line []byte) int {
	x := 0
	for i := range line {
		switch {
		case strings.ContainsRune(vowels, rune(line[i])):
			x += 2
		case strings.ContainsRune(consonants, rune(line[i])):
			for j := i - 1; j > i-4; j-- {
				if j < 0 {
					break
				}
				if strings.ContainsRune(vowels, rune(line[j])) {
					x += 2
					break
				}
			}
			x++
		}
	}
	return x
}
