package main

// "fmt"
// 32 beat var
var h0, h1, h2, h3, h4 uint32 = 0x67452301, 0xEFCDAB89, 0x98BADCFE, 0x10325476, 0xC3D2E1F0

// func for each round
func chooseFunction(i int, b, c, d uint32) uint32 {
	switch {
	case i < 20:
		// функция выбора, ch (choice)
		return (b & c) | ((^b) & d)
	case i < 40:
		return b ^ c ^ d
	case i < 60:
		// функция большинства
		return (b & c) | (b & d) | (c & d)
	default:
		return b ^ c ^ d
	}
}

// func for cost on different round of hashing
func constantFunction(i int) uint32 {
	switch {
	case i < 20:
		return 0x5A827999
	case i < 40:
		return 0x6ED9EBA1
	case i < 60:
		return 0x8F1BBCDC
	default:
		return 0xCA62C1D6
	}
}

// Функция хеширования
// func MyOwnSha(message []byte) [20]byte {

// }

func processBlock(a, b, c, d, e uint32, w [80]uint32) (uint32, uint32, uint32, uint32, uint32) {
	aa, bb, cc, dd := a, b, c, d
	ee := e
	for i := 0; i < 80; i++ {
		temp := leftRotate(aa, 5) + chooseFunction(i, bb, cc, dd) + ee + w[i] + constantFunction(i)

		ee = dd
		dd = cc
		cc = leftRotate(bb, 30)
		bb = aa
		aa = temp
	}
	a = (a + aa) & 0xFFFFFFFF
	b = (b + bb) & 0xFFFFFFFF
	c = (c + cc) & 0xFFFFFFFF
	d = (d + dd) & 0xFFFFFFFF
	e = (e + ee) & 0xFFFFFFFF
	return a, b, c, d, e
}

// Вспомогательная функция для циклического сдвига влево
func leftRotate(value uint32, shiftBits uint32) uint32 {
	return ((value << shiftBits) | (value >> (32 - shiftBits))) & 0xFFFFFFFF
}

func main() {

}
