package belajar_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// implementasi generic interface
// membuat interface yang mengimplementasikan interface
type GetterSetter[T any] interface {
	// memiliki function (sebagai attribute) yang mengimplementasikan generic juga
	GetValue() T
	SetValue(value T)
}

// membuat function yang mengimplementasikan generic
// dan dengan menambahkan parameter dari interface yang generic
func ChangeValue[T any] (param GetterSetter[T], value T) T {
	// mengubah nilai menjadi value yang dikirimkan
	param.SetValue(value)

	// dan mereturn nilai yang baru dengan memanfaatkan function pada GetterSetter
	return param.GetValue()
}

// membuat struct yang mengikuti kontrak interface generic
// maka struct juga harus mengimplementasikan generic seperti interface nya
type MyData[T any] struct {
	// memiliki attribute value
	Value T
}

// membuat method milik struct (mengimplementasikan sesuai kontrak interface generic)
func(d *MyData[T]) GetValue() T {
	return d.Value
}

func(d *MyData[T]) SetValue(value T) {
	d.Value = value
}

// implementasi kode uji generic interface
func TestGenericInterface(t *testing.T) {
	// membuat object my data dari struct
	myData := MyData[string]{} // membuat object dengan attribute kosong

	// mengubah nilai value pada object my data
	result := ChangeValue(&myData, "Taufik")

	// melakukan perbandingan
	assert.Equal(t, "Taufik", result)
	assert.Equal(t, "Taufik", myData.Value)
}
