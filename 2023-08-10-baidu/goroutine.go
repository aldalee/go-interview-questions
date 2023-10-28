// 定义一个切片 arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 然后开 5 个协程并发的修改 arr，修改的逻辑是 arr[i] * 10

package main

import (
	"fmt"
	"sync"
)

// TODO: 下面的代码正确吗
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("arr = ", arr)

	var update = func(int) {
		for i := 0; i < len(arr); i++ {
			arr[i] *= 10
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(id int) {
			defer wg.Done()
			update(id)
			fmt.Printf("goroutine %d: arr = %v\n", id, arr)
		}(i)
	}
	wg.Wait()

	fmt.Println("arr = ", arr)
}
