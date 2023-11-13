package hashing

import (
	"crypto/sha1"
	"reflect"
	"testing"
	// "golang.org/x/text/message"
)

func TestCorrectSha(t *testing.T) {
	// Arrange
	message := "sdfsdfsdfsdf"
	// Заменить местами
	expected := MyOwnSha([]byte(message))
	// Act
	hasher := sha1.New()
	hasher.Write([]byte(message))
	result := hasher.Sum(nil)
	// Assert
	if !reflect.DeepEqual(expected, [20]byte(result)) {
		t.Errorf("Realization of algorithm SHA-1 is wrong. Expect %x, but found %x",
			expected, result)
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
