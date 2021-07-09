package d05

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/**
 * @Author Baofeng Xu
 * @Date 2021/5/13 17:32
 *
 * Description: 纤程
 **/

func TestAtomic() {
	testLoad()
	fmt.Println()
	testAdd()
	fmt.Println()
	testCompereAndSwap()
}

func testLoad() {
	fmt.Println("====================================== atomic > load =====================================")
	var intA int32 = 10
	for {
		value := atomic.LoadInt32(&intA)
		if atomic.CompareAndSwapInt32(&intA, value, value+2) {
			fmt.Println("intA:", intA, ", value:", value)
			break
		}
		fmt.Println("intA:", intA, ", value:", value)
	}
}

func testAdd() {
	fmt.Println("====================================== atomic > add ======================================")
	var intA int32 = 10
	atomic.AddInt32(&intA, 3)
	fmt.Println("intA:", intA)
}

func testCompereAndSwap() {
	fmt.Println("====================================== atomic > compereAndSwap ===========================")
	var intA int32 = 10
	for {
		value := intA
		if atomic.CompareAndSwapInt32(&intA, value, value+1) {
			fmt.Println("intA:", intA, ", value:", value)
			break
		}
		fmt.Println("intA:", intA, ", value:", value)
	}
}

func TestSync() {
	testMutex()
	fmt.Println()
	testRWMutex()
}

var arrB []int
var rwLock sync.RWMutex

func testRWMutex() {
	fmt.Println("====================================== sync > RWMutex =====================================")
	for i := 0; i < 5; i++ {
		go write(i)
	}
	for i := 0; i < 10; i++ {
		go read()
	}
	time.Sleep(20 * time.Second)
	fmt.Println("last:", arrB[0], ", end", arrB[len(arrB)-1], ", length:", len(arrB), ", capacity:", cap(arrB))
}

func read() {
	rwLock.RLock()
	time.Sleep(2 * time.Second)
	fmt.Println("read > arrB:", arrB, ", length:", len(arrB), ", capacity:", cap(arrB))
	rwLock.RUnlock()
}

func write(i int) {
	rwLock.Lock()
	time.Sleep(3 * time.Second)
	arrB = append(arrB, i)
	fmt.Println("write > arrB:", arrB, ", length:", len(arrB), ", capacity:", cap(arrB))
	rwLock.Unlock()
}

var arrA []int
var lock sync.Mutex // 互斥锁

func testMutex() {
	fmt.Println("====================================== sync > mutex =====================================")
	for i := 0; i < 10000; i++ {
		go add(i)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("last:", arrA[0], ", end", arrA[len(arrA)-1], ", length:", len(arrA), ", capacity:", cap(arrA))
}

func add(i int) {
	// ----- 加锁
	lock.Lock()
	// ----- 释放锁
	defer lock.Unlock()
	arrA = append(arrA, i)
}

func TestGoroutine() {
	fmt.Println("====================================== 纤程 ======================================")
	go say("yanWu")
	say("lotus")
}

func say(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(str)
	}
}
