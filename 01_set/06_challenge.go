package main

import (
	"github.com/tomatopeel/pals/datautils"
	"github.com/tomatopeel/pals/futils"
	"log"
	"strings"
)

var (
	remote_file string = "http://cryptopals.com/static/challenge-data/6.txt"
	local_file  string = "secrets_01_06.txt"
	alphabet    string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ \n\r,.'0123456789!()?\";:"
)

func main() {
	data, err := futils.ReadAllBase64File(local_file, remote_file)
	if err != nil {
		log.Fatal(err)
	}

	k := KeySize(1, 50, data)

	blocks := datautils.Blocks(k, data)
	transd := make([][]byte, k)

	for _, block := range blocks {
		for j, byte := range block {
			transd[j] = append(transd[j], byte)
		}
	}

	key := []byte{}
	for _, block := range transd {
		c, _ := TopCharacter(block)
		key = append(key, byte(c))
	}
	log.Printf("KEY: %s", string(key))

	plaintext := datautils.RepeatingKeyXOR(data, key)
	log.Printf("PLAINTEXT: %s", string(plaintext))
}

// Determine most likely keysize between x and y
func KeySize(x, y int, data []byte) (keysize int) {
	var ham float64

	for i := x; i <= y; i++ {

		temp_ham, counter := float64(0), 0
		blocks := datautils.Blocks(i, data)

		for j, init := range blocks {
			for _, rem := range blocks[j+1:] {

				result, err := datautils.Hamming(init, rem)
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

// Determine most likely character as part of key
func TopCharacter(secret []byte) (c rune, score int) {
	tester := make([]byte, len(secret))

	for _, a := range alphabet {
		for i := 0; i < len(secret); i++ {
			tester[i] = byte(a)
		}
		for i := range secret {
			tester[i] ^= secret[i]
		}
		temp_score := Score(tester)
		if temp_score > score {
			score = temp_score
			c = rune(a)
		}
	}
	return
}

func Score(line []byte) (score int) {
	for i := range line {
		if strings.ContainsRune(alphabet, rune(line[i])) {
			score++
		}
	}
	return
}
