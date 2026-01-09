package belajar_golang_generics

import (
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

// implementasi dari experimental package
func ExperimentalMin[T constraints.Ordered](first, second T) T {
	// ordered adalah sebuah type set support operator perbandingan-
	// yang mencakup type parameter dengan tipe data berikut :
	// ~int | ~int8 | ~int16 | ~int32 | ~int64 | (Signed integer)
	// 	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | (Unsigned integer)
	// 	~float32 | ~float64 | (Float)
	// 	~string (String)
	
	// melakukan perbandingan
	if first < second {
		return first
	} else {
		return second
	}
}

// membuat implementasi pengujian yang mengimplementasi function diatas
func TestExperimentalMin(t *testing.T) {
	// melakukan perbandingan dengan assert
	assert.Equal(t, 100, ExperimentalMin(100, 200))
	assert.Equal(t, 100.0, ExperimentalMin(100.0, 200.0))
}

// implementasi experimental maps
func TestExperimentalMaps(t *testing.T) {
	// membuat beberapa data maps
	first := map[string]string {
		"Name": "Taufik",
	}
	second := map[string]string {
		"Name": "Taufik",
	}

	// melakukan perbandingan
	// package maps dan function equal menggunakan generic
	assert.True(t, maps.Equal(first, second)) 
}

// implementasi experimental slices
func TestExperimentalSlices(t *testing.T) {
	// membuat beberapa data slices
	first := []string{"Taufik"}
	second := []string{"Taufik"}

	// melakukan perbandingan
	// package slices dan function equal menggunakan generic
	assert.True(t, slices.Equal(first, second)) 
}
