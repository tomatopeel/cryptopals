package main

import (
	"crypto/aes"
	"encoding/base64"
	"github.com/tomatopeel/pals/futils"
	"io/ioutil"
	"log"
	"os"
)

var (
	key         string = "YELLOW SUBMARINE"
	remote_file string = "http://cryptopals.com/static/challenge-data/7.txt"
	local_file  string = "secrets_01_07.txt"
)

func main() {
	file, err := os.Open(local_file)
	if err != nil {
		futils.DownloadFile(local_file, remote_file)
		file, err = os.Open(local_file)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer file.Close()

	decoder := base64.NewDecoder(base64.StdEncoding, file)
	ciphertext, err := ioutil.ReadAll(decoder)

	result := DecryptAesEcb(ciphertext, []byte(key))
	log.Println(string(result))
}

func DecryptAesEcb(data, key []byte) []byte {
	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}

	decrypted := make([]byte, len(data))
	size := len(key)

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		cipher.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted
}
