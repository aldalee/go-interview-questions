// 分析下面程序的运行结果，并说明原因
package main

import "fmt"

type Fruit interface {
	Eat()
}

type Orange struct{}

func (o *Orange) Eat() {
}

func GetA() *Orange {
	var o *Orange
	return o
}

func GetB() Fruit {
	var o *Orange
	return o
}

func main() {
	fmt.Println(GetA() == nil) // true
	fmt.Println(GetB() == nil) // false
}
