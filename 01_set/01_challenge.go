package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

var (
	orig string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
)

func main() {
	old_main()
	fmt.Println()
	bytes, err := hex.DecodeString(orig)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(bytes))
	fmt.Printf("\nlen(bytes): %d\tcap(bytes): %d\n", len(bytes), cap(bytes))
}

func old_main() {
	result, err := HexStringToString(orig)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(result)))
}

func HexStringToString(in string) (string, error) {
	src := []byte(in)
	dst := make([]byte, hex.DecodedLen(len(src)))
	_, err := hex.Decode(dst, src)
	if err != nil {
		return "", err
	}
	return string(dst), nil
}
