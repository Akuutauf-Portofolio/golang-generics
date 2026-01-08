package belajar_golang_generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// implementasi type parameter
// membuat function dengan type parameter
// T adalah generic parameter (penamannya sengaja tidak lengkap, agar tidak membingungkan dengan parameter lain)
// type parameter 'T' menggunakan type constraint 'interface{}', yang artinya semua tipe data diterima
// interface{} atau alias yang baru bernama 'any'. agar lebih memudahkan

func Length[T any](param T) T {
	fmt.Println(param)
	return param // param adalah variabel di parameter yang berjenis type parameter / T
}

func TestSample(t *testing.T) {
	// mengimplementasikan function Length
	// pada saat pemanggilan function dengan type parameter, harus ditambahkan kurung siku dan di isi tipe data,-
	// disusul dengan value dari tipe data yang di tentukan
	var result string = Length[string]("Taufik") 
	var resultNumber int = Length[int](9999) 
	
	// melakukan perbandingan
	assert.Equal(t, "Taufik", result)
	assert.Equal(t, 9999, resultNumber)
}