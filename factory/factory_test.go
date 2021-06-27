package factory

import (
	"testing"
)

func TestNew(t *testing.T) {
	var (
		student = New(student)
		coder   = New(coder)
	)
	student.Eat()
	student.Sleep()
	coder.Eat()
	coder.Sleep()
}
