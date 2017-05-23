package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	file_url   string = "http://cryptopals.com/static/challenge-data/4.txt"
	local_file string = "/home/tomato/secrets.txt"
	alphabet   string = "abcdefghijklmnopqrstuvwxyz"
	vowels     string = "aeiouAEIOU"
	consonants string = "bcdfghjklmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ"
)

type Line struct {
	original string
	xored    string
	score    int
}

func main() {
	file, err := os.Open(local_file)
	if err != nil {
		file = retrieveFile(local_file, file_url)
	}
	results := ParseFile(file)
	for _, line := range results {
		fmt.Println(line.original)
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
		original := scanner.Text()
		lines[i] = Line{
			original: original,
			xored: Xor(original)
			top_score: Score(
		}
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

func Xor(char int32, secret []byte) string {
	tester := make([]byte, len(secret))
	for i := 0; i < len(secret); i++ {
		tester[i] = byte(char)
	}
	for i := range secret {
		tester[i] ^= secret[i]
	}
	fmt.Printf("%d\t%s\n", Score(tester), string(tester))
	return ""
}

func Score(Line []byte) int {
	x := 0
	for i := range Line {
		switch {
		case strings.ContainsRune(vowels, rune(Line[i])):
			x += 2
		case strings.ContainsRune(consonants, rune(Line[i])):
			for j := i - 1; j > i-4; j-- {
				if j < 0 {
					break
				}
				if strings.ContainsRune(vowels, rune(Line[j])) {
					x += 2
					break
				}
			}
			x++
		}
	}
	return (x * 1000) / len(Line)
}
