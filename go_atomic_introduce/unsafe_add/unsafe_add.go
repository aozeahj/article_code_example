package main

import (
	"fmt"
	"sync"
)

// 在本例子中 count 是一个未初始化的全局变量 存放在.bss段上, 未初始化的全局变量会被初始化为0
// 显然全局 count 是一个共享数据
var count int

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// 协程g1 对全局变量count 进行步长为1的累加操作
	go incrementCount(&wg)

	// 协程g2 对全局变量count 进行步长为1的累加操作
	go incrementCount(&wg)

	// 阻塞等待两个协程执行完毕
	wg.Wait()
	fmt.Println("Final count:", count)
}

func incrementCount(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		count++ // 读取count的值, 然后+1, 然后写回count
	}
}
