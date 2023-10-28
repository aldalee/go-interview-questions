package io

import "fmt"

type Channel struct{}

func (ch *Channel) Scan() {
	fmt.Println("Channel::Scan()")
}

func (ch *Channel) Input() {
	fmt.Println("Channel::Input()")
}

func (ch *Channel) Echo() {
	fmt.Println("Channel::Echo()")
}

func (ch *Channel) Output() {
	fmt.Println("Channel::Output()")
}
