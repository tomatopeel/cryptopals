package main

import (
	"encoding/base64"
	"fmt"
	"github.com/tomatopeel/pals/bitutils"
	"github.com/tomatopeel/pals/futils"
	"io"
	"log"
	"os"
)

var (
	remote_file string = "http://cryptopals.com/static/challenge-data/6.txt"
	local_file  string = "secrets_01_06.txt"
)

func main() {
	f, err := os.Open(local_file)
	if err != nil {
		futils.DownloadFile(local_file, remote_file)
		f, err = os.Open(local_file)
		if err != nil {
			log.Fatal(err)
		}
	}

	d := base64.NewDecoder(base64.StdEncoding, f)

	for i := 2; i <= 40; i++ {
		HamBytes(d, i)
	}
}

func HamBytes(reader io.Reader, k int) int {
	a := make([]byte, k)
	b := make([]byte, k)
	reader.Read(a)
	reader.Read(b)
	result, err := bitutils.Hamming(a, b)
	if err != nil {
		log.Fatal(err)
	}
	result /= k
	fmt.Println(result)
	return result
}
