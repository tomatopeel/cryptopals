package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

var (
	file_url   string = "http://cryptopals.com/static/challenge-data/4.txt"
	local_file string = "secrets_01_04.txt"
	alphabet   string = "abcdefghijklmnopqrstuvwxyz"
	vowels     string = "aeiouAEIOU"
	consonants string = "bcdfghjklmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ"
)

type Line struct {
	Secret   []byte
	TopXored string
	Score    int
}

func main() {
	file, err := os.Open(local_file)
	if err != nil {
		file = retrieveFile(local_file, file_url)
	}
	results := ParseFile(file)
	sort.Slice(results, func(i, j int) bool { return results[i].Score > results[j].Score })

	for i := 0; i < 1; i++ {
		fmt.Printf("%d\n", results[i].Score)
		fmt.Println(string(results[i].TopXored))
	}
}

func ParseFile(file *os.File) []Line {
	scanner := bufio.NewScanner(file)
	n := 0
	for scanner.Scan() {
		n++
	}
	file.Seek(0, 0)

	lines := make([]Line, n)
	scanner = bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line, err := hex.DecodeString(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		lines[i] = Xor(line)
	}

	return lines
}

func retrieveFile(local string, url string) *os.File {
	file, err := os.Create(local)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func Xor(secret []byte) Line {
	line := Line{}
	line.Secret = secret
	tester := make([]byte, len(secret))

	for a := range alphabet {
		for i := 0; i < len(secret); i++ {
			tester[i] = byte(a)
		}
		for i := range secret {
			tester[i] ^= secret[i]
		}
		score := Score(tester)
		if score > line.Score {
			line.Score = score
			line.TopXored = string(tester)
		}
	}

	return line
}

func Score(line []byte) int {
	x := 0
	for i := range line {
		switch {
		case strings.ContainsRune(vowels, rune(line[i])):
			x += 2
		case strings.ContainsRune(consonants, rune(line[i])):
			for j := i - 1; j > i-4; j-- {
				if j < 0 {
					break
				}
				if strings.ContainsRune(vowels, rune(line[j])) {
					x += 2
					break
				}
			}
			x++
		}
	}
	return (x * 1000) / len(line)
}
