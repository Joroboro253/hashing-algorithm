package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"main/hashing"
	"os"
)

func main() {
	photoHashing()
}

func textHashing() {
	var message string
	fmt.Println("Введите сообщение: ")
	fmt.Fscan(os.Stdin, &message)
	// with my realization
	hash := hashing.MyOwnSha([]byte(message))
	fmt.Printf("My own realization of SHA-1 Hash: %x\n", hash)

	// with Library
	hasher := sha1.New()
	hasher.Write([]byte(message))
	bs := hasher.Sum(nil)
	fmt.Printf("SHA-1 Hash with Lib: %x\n", bs)

}

func photoHashing() {
	imagePath := "images/image for hashing.jpg"
	imageBytes, err := ioutil.ReadFile(imagePath)
	fmt.Printf("Message size: %d bytes\n", len(imageBytes))
	// with my realization
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		os.Exit(1)
	}

	// with Library
	hasher := sha1.New()
	hasher.Write([]byte(imageBytes))
	bs := hasher.Sum(nil)
	fmt.Printf("SHA-1 Hash with Lib: %x\n", bs)
}
