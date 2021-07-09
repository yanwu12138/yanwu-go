package d02

import (
	constant "../constant"
	"fmt"
)

/**
 * @Author Baofeng Xu
 * @Date 2021/5/13 10:16
 *
 * Description: 函数和方法
 **/

// TestMethod 方法
func TestMethod() {
	fmt.Println("====================================== 方法 ======================================")
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的半径:", c1.radius, ", 圆的直径:", c1.getDiameter(), ", 圆的面积:", c1.getArea())
}

// Circle /** 结构体
type Circle struct {
	// ----- 半径
	radius float64
}

// getArea /** 计算Circle的面积
func (c Circle) getArea() float64 {
	return constant.PI * c.radius * c.radius
}

// getDiameter /** 计算Circle的直径
func (c Circle) getDiameter() float64 {
	return c.radius * 2
}

// TestClosure 闭包
func TestClosure() {
	fmt.Println("====================================== 闭包 ======================================")
	nextNum := sequence()
	fmt.Println("nextNum1:", nextNum())
	fmt.Println("nextNum2:", nextNum())
	fmt.Println("nextNum3:", nextNum())

	nextNum = sequence()
	fmt.Println("nextNum1:", nextNum())
	fmt.Println("nextNum2:", nextNum())
	fmt.Println("nextNum3:", nextNum())
}

func sequence() func() int {
	result := 0
	return func() int {
		result++
		return result
	}
}

// ReferencePassed 参数引用传递
func ReferencePassed() {
	fmt.Println("====================================== 引用传递 ======================================")
	intA, intB := 100, 200
	fmt.Println("原始数据 > intA:", intA, ", intB:", intB)
	swapReference(&intA, &intB)
	fmt.Println("调用之后 > intA:", intA, ", intB:", intB)
}

func swapReference(intA *int, intB *int) {
	fmt.Println("交换前 > intA:", *intA, ", intB:", *intB)
	*intA, *intB = *intB, *intA
	fmt.Println("交换后 > intA:", *intA, ", intB:", *intB)
}

// ValuePassed 参数值传递
func ValuePassed() {
	fmt.Println("====================================== 值传递 ======================================")
	intA, intB := 100, 200
	fmt.Println("原始数据 > intA:", intA, ", intB:", intB)
	swapValue(intA, intB)
	fmt.Println("调用之后 > intA:", intA, ", intB:", intB)
}

func swapValue(intA, intB int) {
	fmt.Println("交换前 > intA:", intA, ", intB:", intB)
	intA, intB = intB, intA
	fmt.Println("交换后 > intA:", intA, ", intB:", intB)
}

// TestFunc 函数
func TestFunc() {
	fmt.Println("====================================== func ======================================")
	intA, intB := 10, 20
	fmt.Println("原始数据 > intA:", intA, ", intB:", intB)

	intA, intB = swap(intA, intB)
	fmt.Println("进行交换后 > intA:", intA, ", intB:", intB)

	intA, intB = 10, 20
	intA, _ = swap(intA, intB)
	fmt.Println("再次交换，且只接收第一个参数 > intA:", intA, ", intB:", intB)

	intA, intB = 10, 20
	_, intB = swap(intA, intB)
	fmt.Println("再次交换，且只接收第二个参数 > intA:", intA, ", intB:", intB)
}

func swap(intA, intB int) (int, int) {
	return intB, intA
}
