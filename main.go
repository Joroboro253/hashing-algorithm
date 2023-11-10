package main

import (
	"fmt"
)

// "fmt"

var h0, h1, h2, h3, h4 uint32 = 0x67452301, 0xEFCDAB89, 0x98BADCFE, 0x10325476, 0xC3D2E1F0

// func for each round
func f(t uint32, b, c, d uint32) uint32 {
	switch {
	case t < 20:
		// функция выбора, ch (choice)
		return (b & c) | ((^b) & d)
	case t < 40:
		return b ^ c ^ d
	case t < 60:
		// функйция большинства
		return (b & c) | (b & d) | (c & d)
	default:
		return b ^ c ^ d
	}
}

func k(t uint32) uint32 {
	switch {
	case t < 20:
		return 0x5A827999
	case t < 40:
		return 0x6ED9EBA1
	case t < 60:
		return 0x8F1BBCDC
	default:
		return 0xCA62C1D6
	}
}

// Функция хеширования
// func MyOwnSha(message []byte) [20]byte {

// }

func main() {
	message := "Your message here"

	fmt.Println(message)
}
