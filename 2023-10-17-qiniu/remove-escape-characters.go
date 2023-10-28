// 输入：一个带转义字符’\b’的字符串
// 要求：删除转义字符’\b’和它前面的字符，如果遇到多个连续的’\b’，则删除相同数量的转义字符和前面的字符
// 输出：最终的结果字符串
// 举例：输入"abc\b\bd\b\bghi"，期望输出"ghi"；输入"ab\bcd\be\n"，输出 "ace\n"

package main

import "fmt"

func removeBackspace(s string) string {
	var stack []rune
	for _, ch := range s {
		if ch == '\b' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1] // 弹出栈顶元素
			}
		} else {
			stack = append(stack, ch) // 将字符压入栈中
		}
	}
	return string(stack)
}

func main() {
	s := "abc\b\bd\b\bghi"
	fmt.Println(removeBackspace(s)) // Expect: "ghi"

	s = "ab\bcd\be\n"
	fmt.Println(removeBackspace(s)) // Expect: "ace\n"
}
