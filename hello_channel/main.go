package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printer(ch chan int) {
	for i := range ch {
		fmt.Printf("Received %d\n", i)
	}
	wg.Done()
}

func main() {
	c := make(chan int)
	go printer(c)
	wg.Add(1)

	for i := 1; i < 10; i++ {
		// 对通道添加go程，与执行并没有对应关系。在添加后通道根据对应函数进行处理，
		// 与主进程调度无关，有时比主进程后续处理快有时候不是
		c <- i
		fmt.Println("塞入数字", i)
	}

	// 只能保证其与循环添加的先后顺序，实际执行并不能保证。
	fmt.Println("塞入完成")

	// 关闭通道
	close(c)

	// 等待通道内事件结束
	wg.Wait()
}
