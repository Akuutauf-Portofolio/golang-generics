package belajar_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// implementasi in line type constraint
// jadi kita bisa langsung menambahkan constraint tipe data nya secara inline, alih alih menggunakan type set
func FindMin[T interface{int | int64 | float64 }](first, second T) T {
	// melakukan pengecekan
	if first < second {
		return first
	} else {
		return second
	}
} 

// implementasi pengujian dari fungsi diatas
func TestFindMin(t *testing.T) {
	// melakukan perbandingan
	assert.Equal(t, 100, FindMin(100, 100))
	assert.Equal(t, int64(100), FindMin(int64(100), int64(200)))
	assert.Equal(t, 100.0, FindMin(100.0, 200.0))
}

// membuat function yang mengimplementasikan inline
func GetFirst[T []E, E any](slice T) E {
	// mengambil data pertama dari slice E, menggunakan parameter slice / T
	first := slice[0]

	// mengembalikan data pertama
	return first
}

// membuat kode uji untuk mengimplementasikan function diatas
func TestGetFirst(t *testing.T) {
	names := []string {
		"Taufik", "Hidayat",
	}

	// menggunakan function get first
	data := GetFirst[[]string, string](names)

	// melakukan perbandingan
	assert.Equal(t, "Taufik", data)
}