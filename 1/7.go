package main

import (
	"github.com/tomatopeel/cryptopals/crypto/aes"
	"github.com/tomatopeel/cryptopals/futils"
	"log"
)

var (
	key         string = "YELLOW SUBMARINE"
	remote_file string = "http://cryptopals.com/static/challenge-data/7.txt"
	local_file  string = "secrets_01_07.txt"
)

func main() {
	secret, err := futils.ReadAllBase64File(local_file, remote_file)
	if err != nil {
		log.Fatal(err)
	}

	result, err := aes.DecryptEcb(secret, []byte(key))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(result))
}
