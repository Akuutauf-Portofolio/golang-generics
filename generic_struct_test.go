package belajar_golang_generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// implementasi generic struct dan generic method
// membuat type yang menerapkan generic
type Data[T any] struct {
	First, Second T // attribute sama sama menggunakan generic / type parameter
}

// membuat method milik Data yang mengimplementasikan generic
// tidak perlu lagi sebutkan any untuk membuat method generic
// kalau tidak dibutuhkan untuk type parameternya bisa diberikan tana '_' (underscore)
func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

// yang bisa memiliki type parameter itu di function saja, karena dia independent
// kalau method dia nempel / bergantung pada struct, sehingga tidak bisa diberikan type parameter
func (d *Data[T]) ChangeFirst(first T) T {
	// mengubah data first
	d.First = first

	// mereturn data dengan type parameter yang sama
	return d.First
}

// menerapkan kode uji
func TestData(t *testing.T) {
	// membuat object data
	// wajib untuk memasukkan tipe data untuk type parameter nya
	data := Data[string] {
		First: "Taufik",
		Second: "Hidayat",
	}

	fmt.Println(data)
}

// menerapkan kode uji untuk method generic
func TestGenericMethod(t *testing.T) {
	// membuat object data
	data := Data[string] {
		First: "Taufik",
		Second: "Hidayat",
	}

	// melakukan perbandingan dengan assert
	assert.Equal(t, "Budi", data.ChangeFirst("Budi")) // sudah type interfence
	assert.Equal(t, "Hello Budi", data.SayHello("Budi")) // sudah type interfence
}