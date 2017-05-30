package main

import (
	"encoding/base64"
	"github.com/tomatopeel/pals/bitutils"
	"github.com/tomatopeel/pals/futils"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	remote_file string = "http://cryptopals.com/static/challenge-data/6.txt"
	local_file  string = "secrets_01_06.txt"
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
	data, err := ioutil.ReadAll(decoder)
	if err != nil {
		log.Fatal(err)
	}

	k := FindKeySize(1, 50, data)

	blocks := Blocks(k, data)
	transd := make([][]byte, k)

	for _, block := range blocks {
		for j, byte := range block {
			transd[j] = append(transd[j], byte)
		}
	}

	for _, block := range transd {
		log.Printf("%x", block)
	}
}

func FindKeySize(x, y int, data []byte) (keysize int) {
	var ham float64

	for i := x; i <= y; i++ {

		temp_ham, counter := float64(0), 0
		blocks := Blocks(i, data)

		for j, init := range blocks {
			for _, rem := range blocks[j+1:] {

				result, err := bitutils.Hamming(init, rem)
				if err != nil {
					log.Fatal(err)
				}

				temp_ham += (float64(result) / float64(i))
				counter++
			}
		}

		temp_ham /= float64(counter)
		if temp_ham < ham || ham == 0 {
			ham = temp_ham
			keysize = i
		}
	}
	return keysize
}

func Blocks(k int, data []byte) (blocks [][]byte) {
	for b, rem := Block(k, data); b != nil; b, rem = Block(k, rem) {
		blocks = append(blocks, b)
	}
	return
}

func Block(k int, data []byte) ([]byte, []byte) {
	if len(data) < k {
		return nil, nil
	} else {
		return data[:k], data[k:]
	}
}

func Ham(reader io.Reader, keysize int) (float64, error) {
	ham := float64(0)
	a, b := make([]byte, keysize), make([]byte, keysize)
	i := 0
	for {
		n, err := io.ReadFull(reader, a)
		if err != nil && err != io.ErrUnexpectedEOF && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		n, err = io.ReadFull(reader, b)
		if err != nil && err != io.ErrUnexpectedEOF && err != io.EOF {
			log.Fatal(err)
		}
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
