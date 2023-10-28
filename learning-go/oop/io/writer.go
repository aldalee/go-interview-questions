package io

import "fmt"

type Writer struct{}

func (ch *Writer) Echo() {
	fmt.Println("Writer::Echo()")
}

func (ch *Writer) Output() {
	fmt.Println("Writer::Output()")
}
