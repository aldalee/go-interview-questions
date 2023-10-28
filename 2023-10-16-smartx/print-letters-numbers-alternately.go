// 使用两个goroutine交替打印序列，一个goroutine打印数字，另一个goroutine打印字母，最终效果
// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
package main

import (
	"fmt"
	"sync"
)

func main() {
	number := make(chan any)
	letter := make(chan any)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		i := 1
		for range number {
			fmt.Printf("%d%d", i, i+1)
			i += 2
			letter <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		j := 'A'
		for range letter {
			if j > 'Z' {
				close(number)
				close(letter)
				return
			}
			fmt.Printf("%c%c", j, j+1)
			j += 2
			number <- struct{}{}
		}
	}()
	number <- struct{}{}
	wg.Wait()
}
