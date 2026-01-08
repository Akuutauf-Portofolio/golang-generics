package belajar_golang_generics

import (
	"fmt"
	"testing"
)

// implementasi multiple type parameter
func MultipleParameter[T1 any, T2 any](param1 T1, param2 T2) {
	// menampilkan apa yang dikirimkan melalui type parameter
	fmt.Println(param1) // dengan type parameter any
	fmt.Println(param2) // dengan type parameter any
}

func TestMultipleParameter(t *testing.T) {
	// memanggil function multiple parameter, juga wajib mengisikan parameter sesuai jumlah-
	// type parameter yang ada
	// type parameter yang sudah ditentukan tipe datanya, tidak boleh berbeda value nya
	MultipleParameter[string, int]("Taufik", 100) 
	MultipleParameter[int, string](9999, "Hidayat") 
}