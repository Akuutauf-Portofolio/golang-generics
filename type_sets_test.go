package belajar_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// implementasi dari type sets, type delcaration, type approximation dan type_inference_test

// membuat type declaration Age untuk alias int
type Age int

// membuat kumpulan tipe data yang digunakan sebagai type parameter dengan menyiapkan type sets
type Number interface {
	// tipe data yang sudah di definisikan di sini akan dapat digunakan sebagai type parameter 
	~int | int8 | int16 | int32 | int64 | float32 | float64

	// type approximation
	// jika diberikan tanda '~' (tilde), diawal tipe data, maka otomatis semua tipe data dengan alias/-
	// type declaration yang memiliki tipe data serupa (contoh nya int, dengan Age), maka akan bisa digunakan di type parameter
}

// kemudian type sets diatas di daftarkan sebagai constraint di function ini (sebagai type parameter)
func Min[T Number](first, second T) T {
	// karena type sets nya bida dilakukan perbandingan, maka kita bisa memberikan contoh logika perbandingan
	if first < second {
		return first
	} else {
		return second
	}
}

// menerapkan type parameter dan type sets ke dalam pengujian
func TestMin(t *testing.T) {
	// melakukan perbandingan dengan assert
	assert.Equal(t, int(100), Min[int](100, 200))
	assert.Equal(t, int64(100), Min[int64](100, 200))
	assert.Equal(t, float64(100.0), Min[float64](100.0, 200.0))
	
	// namun ketika type declaration yang sudah kita buat sebelumnya (Age) kita implementasikan funsi min-
	// maka hal tersebut tidak bisa, karea meski secara harfiah sama sama int, namun di type sets Age tidak terdfatar
	// harus ditambahkan terlebih dahulu di type sets

	// atau cukup menggunakan operator type approximation
	assert.Equal(t, Age(100), Min[Age](100, 200)) 
}

func TestTypeInference(t *testing.T) {
	// dengan menggunakan type inference, kita bisa langsung mengisikan nilai, tanpa memberikan type parameternya

	// melakukan perbandingan dengan assert
	assert.Equal(t, int(100), Min(100, 200))
	assert.Equal(t, int64(100), Min(int64(100), int64(200))) // karena sama dengan value tipe data int, maka dilakukan konversi 
	assert.Equal(t, float64(100.0), Min(100.0, 200.0))
	assert.Equal(t, Age(100), Min(Age(100), Age(200))) 
}