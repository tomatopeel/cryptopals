package main

import (
	"fmt"
	"github.com/tomatopeel/pals/bitutil"
	"log"
)

var (
	remote_file string = "http://cryptopals.com/static/challenge-data/6.txt"
	local_file  string = "secrets_01_06.txt"
	ham_a       string = "this is a test"
	ham_b       string = "wokka wokka!!!"
)

func main() {
	result, err := bitutil.Hamming([]byte(ham_a), []byte(ham_b))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
