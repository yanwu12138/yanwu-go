package d01

import (
	"fmt"
)

/**
 * @Author Baofeng Xu
 * @Date 2021/5/13 9:56
 *
 * Description: 运算符
 **/

// OtherOperators 其他运算符
func OtherOperators() {
	fmt.Println("====================================== 其他运算符 ======================================")
	var intA = 4
	var intB int32
	var intC float32
	var ptrA *int
	fmt.Println("intA:", intA, ", intB:", intB, ", intC:", intC, ", ptrA:", ptrA)

	// ----- & 返回变量存储的地址
	// ----- * 指针变量
	ptrA = &intA
	fmt.Println("ptrA:", ptrA, ", *ptrA:", *ptrA)
}

// AssignmentOperator 赋值运算符
func AssignmentOperator() {
	fmt.Println("====================================== 赋值运算符 ======================================")
	intA := 21
	var intC int
	fmt.Println("intA:", intA, ", intC:", intC)

	intC = intA
	fmt.Println("intA:", intA, ", intC:", intC)

	intC += intA
	fmt.Println("intA:", intA, ", intC:", intC)

	intC -= intA
	fmt.Println("intA:", intA, ", intC:", intC)

	intC *= intA
	fmt.Println("intA:", intA, ", intC:", intC)

	intC /= intA
	fmt.Println("intA:", intA, ", intC:", intC)

	intC = 200

	intC <<= 2
	fmt.Println("intC <<= 2:", intC)

	intC >>= 2
	fmt.Println("intC >>= 2:", intC)

	intC &= 2
	fmt.Println("intC &= 2:", intC)

	intC |= 2
	fmt.Println("intC |= 2:", intC)

	intC ^= 2
	fmt.Println("intC ^= 2:", intC)
}

// BitwiseOperator 位运算符
func BitwiseOperator() {
	fmt.Println("====================================== 位运算符 ========================================")
	var intA, intB uint = 60, 13
	fmt.Println("与: intA & intB >", intA&intB)
	fmt.Println("或: intA | intB >", intA|intB)
	fmt.Println("异或: intA ^ intB >", intA^intB)
	fmt.Println("左移: intA << 2 >", intA<<2)
	fmt.Println("右移: intA >> 2 >", intA>>2)
}

// LogicalOperator 逻辑运算符
func LogicalOperator() {
	fmt.Println("====================================== 逻辑运算符 ======================================")
	boolA, boolB := true, false
	fmt.Println("与: a && b >", boolA && boolB)
	fmt.Println("或: a || b >", boolA || boolB)
	fmt.Println("非: !a >", !boolA, ", !b >", !boolB)
}

// RelationalOperator 关系运算符
func RelationalOperator() {
	fmt.Println("====================================== 关系运算符 ======================================")
	intA, intB := 10, 100
	fmt.Println("等等:", intA == intB)
	fmt.Println("不等:", intA != intB)
	fmt.Println("大于:", intA > intB)
	fmt.Println("小于:", intA < intB)
	fmt.Println("大于等于:", intA >= intB)
	fmt.Println("小于等于:", intA <= intB)
}

// ArithmeticOperator 算数运算符
func ArithmeticOperator() {
	fmt.Println("====================================== 算术运算符 ======================================")
	intA, intB, intC, intD := 10, 100, 1000, 4
	fmt.Println("加法:", add(intA, intB))
	fmt.Println("减法:", sub(intC, intA))
	fmt.Println("乘法:", multiply(intB, intC))
	fmt.Println("除法:", div(intC, intA))
	fmt.Println("取模:", mod(intA, intD))
	intA++
	fmt.Println("自增:", intA)
	intA--
	fmt.Println("自减:", intA)
}

/**
 * 加法
 */
func add(int1, int2 int) int {
	return int1 + int2
}

/**
 * 减法
 */
func sub(int1, int2 int) int {
	return int1 - int2
}

/**
 * 乘法
 */
func multiply(int1, int2 int) int {
	return int1 * int2
}

/**
 * 除法
 */
func div(int1, int2 int) int {
	return int1 / int2
}

/**
 * 取模
 */
func mod(int1, int2 int) int {
	return int1 % int2
}
