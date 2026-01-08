package belajar_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// implementasi comparable
// tipe data comparable itu sama seperti any / interface{}
// namun tipe data ini memeliki kelebihan yaitu, mampu membandingkan suatu nilai,-
// dengan operator perbandingan != dan ==
func IsSame[T comparable](value1, value2 T) bool {
	// melakukan pengecekan type parameter
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestIsSame(t *testing.T) {
	assert.True(t, IsSame[string]("Taufik", "Taufik"))
	assert.True(t, IsSame[int](1000, 1000))
	assert.True(t, IsSame[bool](true, true))
}