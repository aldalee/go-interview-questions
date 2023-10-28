package main

import "fmt"

func interfaceIsNil(x interface{}) {
	if x != nil {
		fmt.Println("non-empty interface")
		return
	}
	fmt.Println("empty interface")
}

func main() {
	var x interface{} = nil
	var y *int = nil

	fmt.Println(x == nil) // true
	fmt.Println(y == nil) // true

	interfaceIsNil(x) // empty interface
	interfaceIsNil(y) // non-empty interface
}
