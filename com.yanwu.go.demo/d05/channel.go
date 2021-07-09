package d05

import (
	"fmt"
	"time"
)

/**
 * @Author Baofeng Xu
 * @Date 2021/5/13 18:29
 *
 * Description: 通道
 **/

func TestTimeout() {
	fmt.Println("====================================== channel > select > timeout ===========================")
	chanA := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		chanA <- "result 1"
	}()
	select {
	// ----- 从chanA中读取数据
	case res := <-chanA:
		fmt.Println(res)
	// ----- 超时处理
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
}

func TestSelect() {
	fmt.Println("====================================== channel > select ===========================")
	chanA := make(chan int)
	chanB := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("i:", <-chanA)
		}
		chanB <- 0
	}()
	fibonacci(chanA, chanB)
}

func fibonacci(chanA, chanB chan int) {
	x, y := 0, 1
	for {
		select {
		// ----- 将x发送到chanA
		case chanA <- x:
			x, y = y, x+y
		// ----- 当能从chanB中读取数据时，退出操作
		case <-chanB:
			fmt.Println("quit")
			return
		}
	}
}

func TestChanClose() {
	fmt.Println("====================================== channel > close ===========================")
	chanA := make(chan int, 10)
	go testClose(cap(chanA), chanA)
	for i := range chanA {
		fmt.Println("i:", i)
	}
}

func testClose(intA int, chanA chan int) {
	x, y := 0, 1
	for i := 0; i < intA; i++ {
		chanA <- x
		x, y = y, x+y
	}
	// ----- 关闭通道
	close(chanA)
}

func TestChanBuffer() {
	fmt.Println("====================================== channel > buffer ===========================")
	// ----- 定义一个缓冲区大小为2的通道
	chanA := make(chan int, 2)
	// ----- 因为 ch 是带缓冲的通道，我们可以同时发送两个数据, 而不用立刻需要去同步读取数据
	chanA <- 100
	chanA <- 200
	// ----- 读取两个数据
	intA, intB := <-chanA, <-chanA
	fmt.Println("intA:", intA, ", intB:", intB)
}

func TestChannel() {
	fmt.Println("====================================== channel ===========================")
	arr := []int{7, 2, 8, -9, 4, 0}
	chanA := make(chan int)
	go sum(arr[:len(arr)/2], chanA)
	go sum(arr[len(arr)/2:], chanA)
	r, l := <-chanA, <-chanA
	fmt.Println("r:", r, ", l:", l, ", sum:", r+l)
}

func sum(arr []int, chanA chan int) {
	sum := 0
	for _, item := range arr {
		sum += item
	}
	// ----- 把 sum 发送到通道 c
	chanA <- sum
}
