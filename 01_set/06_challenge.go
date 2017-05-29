package main

import (
	"encoding/base64"
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
	defer f.Close()

	keysize := FindKeySize(1, 50, f)
	log.Println(keysize)
}

func FindKeySize(x, y int, f *os.File) int {
	var (
		ham     float64
		keysize int
	)
	for i := x; i <= y; i++ {
		decoder := base64.NewDecoder(base64.StdEncoding, f)
		temp_ham, err := Ham(decoder, i)
		if err != nil {
			log.Fatal(err)
		}
		if temp_ham < ham || ham == 0 {
			ham = temp_ham
			keysize = i
		}
		f.Seek(0, 0)
	}
	return keysize
}

func Ham(reader io.Reader, keysize int) (float64, error) {
	ham := float64(0)
	a, b := make([]byte, keysize), make([]byte, keysize)
	i := 0
	for {
		n, err := reader.Read(a)
		if n == 0 {
			break
		}
		n, err = reader.Read(b)
		if n == 0 {
			break
		}
		result, err := bitutils.Hamming(a, b)
		normalised := float64(result) / float64(keysize)
		if err != nil {
			log.Fatal(err)
		}
		ham += float64(normalised)
		i++
	}
	return ham / float64(i), nil
}
