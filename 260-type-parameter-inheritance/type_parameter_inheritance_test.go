package _60_type_parameter_inheritance

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](parameter T) string {
	return parameter.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePresident struct {
	Name string
}

func (m *MyVicePresident) GetName() string {
	return m.Name
}

func (m *MyVicePresident) GetVicePresidentName() string {
	return m.Name
}

func TestTypeParameterInheritance(t *testing.T) {
	assert.Equal(t, "Eko", GetName[Employee](&MyManager{Name: "Eko"}))
	assert.Equal(t, "Eko", GetName[VicePresident](&MyVicePresident{Name: "Eko"}))
}
