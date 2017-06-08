package main

import (
	"github.com/tomatopeel/cryptopals/crypto/aes"
	"github.com/tomatopeel/cryptopals/futils"
	"log"
)

var (
	remote_file string = "http://cryptopals.com/static/challenge-data/10.txt"
	local_file  string = "secrets_02_10.txt"
	key         string = "YELLOW SUBMARINE"
	blockSize   int    = 16
)

func main() {
	ciphertext, err := futils.ReadAllBase64File(local_file, remote_file)
	if err != nil {
		log.Fatal(err)
	}

	iv := make([]byte, blockSize)
	result, err := aes.DecryptCbc(ciphertext, []byte(key), iv)
	if err != nil {
		log.Fatal(err)
	}

	iv = make([]byte, blockSize)
	result, err = aes.EncryptCbc(result, []byte(key), iv)
	if err != nil {
		log.Fatal(err)
	}

	iv = make([]byte, blockSize)
	result, err = aes.DecryptCbc(result, []byte(key), iv)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(result))

}
