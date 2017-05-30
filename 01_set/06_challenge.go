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
	blocks := Blocks(keysize, f)
	transd := make([][]byte, keysize)

	for _, block := range blocks {
		for j, byte := range block {
			transd[j] = append(transd[j], byte)
		}
	}
}

func Blocks(keysize int, f *os.File) (blocks [][]byte) {
	decoder := base64.NewDecoder(base64.StdEncoding, f)

	for {
		block := make([]byte, keysize)
		n, err := decoder.Read(block)
		if n == 0 {
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		blocks = append(blocks, block)
	}
	return
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

		n, err := f.Seek(0, 0)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("seek's n: %d", n)
	}
	return keysize
}

func Ham(reader io.Reader, keysize int) (float64, error) {
	ham := float64(0)
	a, b := make([]byte, keysize), make([]byte, keysize)
	i := 0
	for {
		n, err := reader.Read(a)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		log.Printf("a: %x\tn=%d\tkeysize=%d", a, n, keysize)
		if n != keysize {
			break
		}
		n, err = reader.Read(b)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		log.Printf("b: %x\tn=%d\tkeysize=%d", b, n, keysize)
		if n != keysize {
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
	log.Println("and out...")
	return ham / float64(i), nil
}
