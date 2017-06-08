package main

import (
	"log"
)

var (
	plaintext string = "YELLOW SUBMARINE"
)

func main() {
	newBlock := Pad([]byte(plaintext), 20)
	log.Printf("% x", newBlock)
	log.Println(string(newBlock))
}

func Pad(block []byte, blockLength int) []byte {
	if len(block) < blockLength {
		newBlock := make([]byte, blockLength)
		i := copy(newBlock, block)
		for ; i < blockLength; i++ {
			newBlock[i] = byte(4)
		}
		return newBlock
	}
	return block
}
