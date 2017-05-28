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

type Blocks [][]byte

func main() {
	f, err := os.Open(local_file)
	if err != nil {
		futils.DownloadFile(local_file, remote_file)
		f, err = os.Open(local_file)
		if err != nil {
			log.Fatal(err)
		}
	}

	decoder := base64.NewDecoder(base64.StdEncoding, f)
	for i := 2; i <= 40; i++ {
		fmt.Println(Ham(decoder, 6, i))
	}
}

func Ham(reader io.Reader, keysize, n_blocks int) float64 {
	ham := float64(0)
	blocks := make([][]byte, n_blocks)
	for i := range blocks {
		blocks[i] = make([]byte, keysize)
		_, err := reader.Read(blocks[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	for i := range blocks {
		for j := i + 1; j < len(blocks); j++ {
			result, err := bitutils.Hamming(blocks[i], blocks[j])
			if err != nil {
				log.Fatal(err)
			}
			ham += float64(result / keysize)
		}
	}
	return ham / float64(n_blocks)
}
