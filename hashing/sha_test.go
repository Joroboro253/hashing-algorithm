package hashing

import (
	"crypto/sha1"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

// Тест на правильность работы релизации алгоритма в сравнении с библиотечной реализацией
func TestCorrectSha(t *testing.T) {
	message := "Hello world"
	result := MyOwnSha([]byte(message))
	hasher := sha1.New()
	hasher.Write([]byte(message))
	expected := hasher.Sum(nil)
	// Assert
	if reflect.DeepEqual([20]byte(result), expected) {
		t.Errorf("Realization of algorithm SHA-1 is wrong. Expect %x, but found %x",
			result, expected)
	}
}

// Бенчмарк демонстрирующий времени работы моей реализации алгоритма
func BenchmarkFunction1(b *testing.B) {
	message := []byte("3464364363")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		MyOwnSha(message)
	}
}

// Бенчмарк демонстрирующий времени работы алгоритма из библиотеки
func BenchmarkFunction2(b *testing.B) {
	message := []byte("3464364363")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		hasher := sha1.New()
		hasher.Write([]byte(message))
		hasher.Sum(nil)
	}
}

// Бенчмарк демонстрирующий время работы моей реализации алгоритма для изображения (состоит из нескольких блоков)
func BenchmarkFunctionPhotoHashing2(b *testing.B) {
	imagePath := "../images/image for hashing.jpg"
	imageBytes, err := ioutil.ReadFile(imagePath)
	if err != nil {
		os.Exit(1)
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		MyOwnSha([]byte(imageBytes))
	}
}

// Бенчмарк демонстрирующий времени работы алгоритма из библиотеки для изображения (состоит из нескольких блоков)
func BenchmarkFunctionPhotoHashing(b *testing.B) {
	imagePath := "../images/image for hashing.jpg"
	imageBytes, err := ioutil.ReadFile(imagePath)
	if err != nil {
		os.Exit(1)
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		hasher := sha1.New()
		hasher.Write([]byte(imageBytes))
		hasher.Sum(nil)
	}
}
