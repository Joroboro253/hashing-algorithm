package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	const input = "This is just a simple test."
	h := sha256.New()
	h.Write([]byte(input))

	fmt.Printf("%x\n", h.Sum(nil))
}
