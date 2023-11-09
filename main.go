package main

import (
	"crypto/sha256"
	"fmt"
)

// Function for padding
func padding(message []byte) []byte {

}

// Permutation function
func permutation(state [][]uint64, outputLength int) []byte {

}

// SHA3 func
func sha3(message []byte, outputLength int) []byte {

}

func main() {
	const input = "This is just a simple test."
	h := sha256.New()
	h.Write([]byte(input))

	fmt.Printf("%x\n", h.Sum(nil))
}
