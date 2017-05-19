package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

var (
	firstStr  string = "1c0111001f010100061a024b53535009181c"
	secondStr string = "686974207468652062756c6c277320657965"
)

func main() {
	first, err := hex.DecodeString(firstStr)
	second, err := hex.DecodeString(secondStr)
	if len(first) != len(second) {
		log.Fatal("Length mismatch")
	}
	if err != nil {
		log.Fatal(err)
	}
	for i := range first {
		first[i] ^= second[i]
	}
	fmt.Println(hex.EncodeToString(first))
}
