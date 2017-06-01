package main

import (
	"encoding/hex"
	"github.com/tomatopeel/pals/datautils"
	"github.com/tomatopeel/pals/futils"
	"log"
)

var (
	localFile  string = "cipher_01_08.txt"
	remoteFile string = "http://cryptopals.com/static/challenge-data/8.txt"
)

func main() {
	data, err := futils.ReadLinesHexFile(localFile, remoteFile)
	if err != nil {
		log.Fatal(err)
	}

	hams := make(map[string]float64)
	for _, block := range data {
		blocks := datautils.Blocks(16, block)

		ham := float64(0)
		counter := 0

		for i, init := range blocks {
			for _, rem := range blocks[i+1:] {
				result, err := datautils.Hamming(init, rem)
				if err != nil {
					log.Fatal(err)
				}
				ham += float64(result)
				counter++
			}
		}

		ham /= float64(counter)

		hams[hex.EncodeToString(block)] = ham
	}

	for k, v := range hams {
		if v < float64(60) {
			log.Printf("it is: %s", k)
		}
	}
}
