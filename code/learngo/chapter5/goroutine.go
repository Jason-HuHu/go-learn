package main

import (
	"fmt"
	"time"
)

func main() {
	c = make(chan int)
	go ready("Tea",2)
	go ready("Coffee",1)
	fmt.Println("I'm waiting,but not too long ")
	//time.Sleep(5 * time.Second)
	//<- c
	//<- c
	i := 0
	L:for{
			select {
			case <- c:
				i++
				if i > 1 {
					break L
				}
			}
		}
	}

/**
	Go 使用channel和goroutine开发并行程序

	为什么叫goroutine？因为已有的短语——线程、协程、进程等都传递不了准确的含义。goroutine的简单模型：它是与其他goroutine并行执行的，有着相同地址空间的函数。它是轻量的，仅仅比分配栈空间多一点点消耗

	普通函数调用：ready("Tea",2)
	作为goroutine运行：go ready("Tea",2)

	ci <- 1 发送整数1 到channel ci
	<- ci 从channel ci 接收整数
	i := <-ci 从channel ci接收整数，并保存到整数i中

	虽然goroutine 是并发执行的，但是它们并不是并行执行的，如果不告诉Go额外的东西，同一时刻只会有一个goroutine执行，利用runtime.GOMAXPROCESS(n)
	可以设置goroutine并行执行的最大数量

 */

//Go routine实践
var c chan int
func ready(w string,sec int ) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w,"is ready")
	c <- 1
}