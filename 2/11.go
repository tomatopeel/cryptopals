package main

import (
	"crypto/rand"
	"fmt"
	"github.com/tomatopeel/cryptopals/crypto/aes"
	"log"
)

func main() {
	_, err := EncryptionOracle([]byte("YELLOW SUBMARINES R US 4 EVA LOLOLOLOL"))
	if err != nil {
		log.Fatal(err)
	}
}

func EncryptionOracle(input []byte) ([]byte, error) {
	key := make([]byte, 16)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatal(err)
	}

	iv := make([]byte, 16)

	blurb, err := aes.EncryptCbc(input, key, iv)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(blurb))

	dec, err := aes.DecryptCbc(blurb, key, iv)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(dec))
	return blurb, nil
}
