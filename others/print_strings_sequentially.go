// 定义两个goroutine顺序的打印"hello, world"，效果如下：
// p1 = h
// p2 = e
// p1 = l
// p2 = l
// p1 = o
// p2 = ,
// p1 =
// p2 = w
// p1 = o
// p2 = r
// p1 = l
// p2 = d

package main

import (
	"fmt"
	"sync"
)

func main() {
	str := "hello, world"
	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < len(str); i += 2 {
			<-ch1
			fmt.Println("p1 =", string(str[i]))
			ch2 <- struct{}{}
		}
		close(ch2)
	}()

	go func() {
		defer wg.Done()
		for i := 1; i < len(str); i += 2 {
			<-ch2
			fmt.Println("p2 =", string(str[i]))
			ch1 <- struct{}{}
		}
		close(ch1)
	}()

	ch1 <- struct{}{}
	wg.Wait()
}
