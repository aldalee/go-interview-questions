// 分析下面程序的运行结果，并说明原因
package main

import "fmt"

func main() {
	s := []int{1}
	fmt.Println(cap(s)) // 1

	s = append(s, 2)
	fmt.Println(cap(s)) // 2

	s = append(s, 3)
	fmt.Println(cap(s)) // 4

	x := append(s, 4)
	fmt.Println(cap(s)) // 4

	y := append(s, 5)
	fmt.Println(cap(s)) // 4

	fmt.Println(s, x, y) //[1 2 3] [1 2 3 5] [1 2 3 5]
}
