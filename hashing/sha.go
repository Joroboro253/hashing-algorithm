package hashing

import (
	"encoding/binary"
)

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
func MyOwnSha(message []byte) [20]byte {
	originalLength := uint64(len(message) * 8)
	message = append(message, 0x80) // Adding 1 bit
	// Дополняем сообщение нулями, если оно не кратно размеру блока
	for len(message)%64 != 56 {
		message = append(message, 0x00)
	}
	// Добавляем длину исходного сообщения в битах в конец сообщения
	lengthBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(lengthBytes, originalLength)
	message = append(message, lengthBytes...)
	// Инициализируем хеш-значения для этого блока
	var a, b, c, d, e uint32 = h0, h1, h2, h3, h4
	// Обработка каждого блока
	for blockStart := 0; blockStart < len(message); blockStart += 64 {
		var w [80]uint32
		for i := 0; i < 16; i++ {
			w[i] = binary.BigEndian.Uint32(message[blockStart+4*i:])
		}

		for i := 16; i < 80; i++ {
			w[i] = leftRotate(w[i-3]^w[i-8]^w[i-14]^w[i-16], 1)
		}

		a, b, c, d, e = processBlock(a, b, c, d, e, w)
	}
	// Формируем итоговое хеш-значение
	var digest [20]byte
	binary.BigEndian.PutUint32(digest[0:4], a)
	binary.BigEndian.PutUint32(digest[4:8], b)
	binary.BigEndian.PutUint32(digest[8:12], c)
	binary.BigEndian.PutUint32(digest[12:16], d)
	binary.BigEndian.PutUint32(digest[16:20], e)

	return digest
}

// input transformation
func processBlock(a, b, c, d, e uint32, w [80]uint32) (uint32, uint32, uint32, uint32, uint32) {
	aa, bb, cc, dd, ee := a, b, c, d, e
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
