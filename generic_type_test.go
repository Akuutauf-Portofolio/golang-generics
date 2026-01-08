package belajar_golang_generics

import (
	"fmt"
	"testing"
)

// implementasi generic type

// membuat type declaration dalam bentuk generic
type Bag[T any] []T // membuat type dengan tipe parameternya any, dan return nya []T

// membuat function dalam bentuk generic yang m	enerapkan type Bag
func PrintBag[T any](bag Bag[T]) {
	// karena bag nya adalah slice, kita bisa print dengan melakukan perulangan
	for _, value := range bag {
		fmt.Println(value)
	}
}

// mengimplementasikan pengujiannya
func TestBagString(t *testing.T) {
	names := Bag[string]{"Taufik", "Ilham", "Dimas"}

	PrintBag(names)
}

func TestBagInt(t *testing.T) {
	numbers := Bag[int]{1, 2, 3, 4, 5}

	PrintBag[int](numbers) // bebas bisa di deklarasikan atau tidak untuk type parameter nya
}