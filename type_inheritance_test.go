package belajar_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// implementasi type parameter inheritance
// membuat interface baru
type Employee interface {
	// yang memiliki method GetName
	GetName() string
}

// membuat function yang mengimplementasikan interface employee
func GetName[T Employee](parameter T) string {
	// kalau kita mengimplementasikan interface yang memiliki method di dalam function-
	// yang mengimplementasikan type parameter, maka kita bisa memanggil method yang ada-
	// di interface tersebut secara langsung melalui type parameter / T
	return parameter.GetName()
}

// membuat interfae baru
type Manager interface {
	GetName() string // sama seperti kontrak pada interface employe
	GetManagerName() string
}

// membuat struct
type MyManager struct {
	Name string
}

// membuat method untuk struct my manager dengan mengimplementasikan kontrak interface Manager
// kalau implementasi untuk method struct, harus menggunakan pointer
// meskipun di golang tidak ada inheritance, namun dari sini terlihat bahwasannya-
// MyManager mengimplementasikan kontrak dari Manager, dan Manager mengimplementasikan kontrak dari Employee
func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

// membuat interface baru lagi
type VicePresident interface {
	GetName() string // mengimplementasikan kontrak dari Employee
	GetVicePresidentName() string
}

// membuat struct lagi
type MyVicePresident struct {
	Name string
}

// implement method untuk my vice president
func (m *MyVicePresident) GetName() string {
	return m.Name
}

func (m *MyVicePresident) GetVicePresidentName() string {
	return m.Name
}

// menguji inheritance
func TestGetName(t *testing.T) {
	// melakukan perbandingan dengan assert
	assert.Equal(t, "Taufik", GetName[Manager](&MyManager{Name: "Taufik"}))
	assert.Equal(t, "Taufik", GetName[VicePresident](&MyVicePresident{Name: "Taufik"}))
}