package main

import (
	"crypto/sha1"
	"fmt"
	"main/hashing"
)

func main() {
	// message := "Your message here"
	message := "Your message here"

	hash := hashing.MyOwnSha([]byte(message))
	fmt.Printf("SHA-1 Hash: %x\n", hash)

	// Lets test with Library
	hasher := sha1.New()
	hasher.Write([]byte(message))
	bs := hasher.Sum(nil)
	fmt.Println(message)
	fmt.Printf("SHA-1 Hash with Lib: %x\n", bs)
	fmt.Println(message)
}
