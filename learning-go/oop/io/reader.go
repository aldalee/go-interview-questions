package io

import "fmt"

type Reader struct{}

func (r *Reader) Scan() {
	fmt.Println("Reader::Scan()")
}

func (r *Reader) Input() {
	fmt.Println("Reader::Input()")
}
