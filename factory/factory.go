package factory

import (
	"log"
)

// IFactory 这是一个接口
type IFactory interface {
	Eat()
	Sleep()
}

type Type string

const (
	student Type = "student"
	coder   Type = "coder"
)

func New(mode Type) IFactory {
	switch mode {
	case student:
		return NewAStudent()
	case coder:
		return NewACoder()
	default:
		return NewACoder()
	}
}

// AStudent a student
type AStudent struct{}

func NewAStudent() IFactory {
	return &AStudent{}
}

func (a *AStudent) Eat() {
	log.Println("eat eat eat")
}

func (a *AStudent) Sleep() {
	log.Println("sleep!!!")
}

// ACoder a coder
type ACoder struct{}

func NewACoder() IFactory {
	return &ACoder{}
}

func (a *ACoder) Eat() {
	log.Println("eat")
}

func (a *ACoder) Sleep() {
	log.Println("sleep sleep!!")
}
