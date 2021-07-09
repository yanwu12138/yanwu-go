package d03

import "fmt"

/**
 * @Author Baofeng Xu
 * @Date 2021/5/13 13:57
 *
 * Description: 指针
 **/

func TestPointer() {
	fmt.Println("====================================== 指针 ======================================")
	// ----- 声明变量
	var intA = 20
	// ----- 声明指针变量

	ptrA := &intA
	fmt.Println("intA:", intA, ", 变量的地址:", &intA, ", 指针的地址:", ptrA, ", 指针的地址值:", *ptrA)

	var ptrB *int
	fmt.Println("空指针:", ptrB)

	// ----- 指针数组，该数组中存储的每个元素都是地址
	arr := [...]int{1, 2, 3}
	var ptrArr [len(arr)]*int
	fmt.Println("数组:", arr, ", 指针数组:", ptrArr)
	ptrArr[0] = &arr[0]
	ptrArr[1] = &arr[1]
	ptrArr[2] = &arr[2]
	fmt.Println("数组:", arr, ", 指针数组:", ptrArr)

	// ----- 指向指针的指针
	intB := 10
	ptrB1 := &intB
	ptrB2 := &ptrB1
	fmt.Println("intB:", intB, ", ptrB1:", ptrB1, ", *ptrB1:", *ptrB1, ", ptrB2", ptrB2, ", *ptrB2", *ptrB2)

	// ----- 将指针作为参数传递（引用传递）
	intC, intD := 10, 20
	fmt.Println("原始数据 > intC:", intC, ", intD:", intD)
	swap(&intC, &intD)
	fmt.Println("调用之后 > intC:", intC, ", intB:", intD)
}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}
