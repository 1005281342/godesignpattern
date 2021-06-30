package strategy

import "fmt"

type Log struct {
	ILogging
}

type ILogging interface {
	Info()
	Error()
}

type DBLog struct{}

func (d DBLog) Error() {
	fmt.Println("Error db log")
}
func (d DBLog) Info() {
	fmt.Println("Info db log")
}

type FileLog struct{}

func (f FileLog) Error() {
	fmt.Println("Error file log")
}
func (f FileLog) Info() {
	fmt.Println("Info file log")
}
