// 定义两个goroutine，交替打印producer和consumer10次
package main

import (
	"fmt"
	"sync"
)

func main() {
	producer := make(chan any, 1)
	consumer := make(chan any, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-producer
			fmt.Println("Producer")
			consumer <- struct{}{}
		}
		close(consumer)
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-consumer
			fmt.Println("Consumer")
			producer <- struct{}{}
		}
		close(producer)
	}()
	producer <- struct{}{}
	wg.Wait()
}
