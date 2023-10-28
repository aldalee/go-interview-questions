package main

import "fmt"

func B() {
	j := A() // position 1
	fmt.Println(j)
}

func A() int {
	i := 1
	defer func() {
		i = 2 // position 2
	}()
	i = 3
	return i // position 3
}

func main() {
	B() // 3

	s1 := []int{1, 2, 3}
	s2 := s1
	s2 = append(s2, 4)

	fmt.Printf("s1 = %v, len = %d, cap = %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2 = %v, len = %d, cap = %d\n", s2, len(s2), cap(s2))

	s2[1] = 100
	fmt.Println(s1[1]) // 2 why?

	fmt.Println(s1[3]) // panic: runtime error: index out of range [3] with length 3
}
