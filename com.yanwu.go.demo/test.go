// ----- 包声明
package main

// ----- 导包
import (
	"./d01"
	"./d02"
	"./d03"
	"./d04"
	"./d05"
	"fmt"
	"math"
)

// ----- 函数
func main() {
	// ----- 变量
	str := "hello,world!"
	/*
	 *	语句
	 */
	fmt.Println(str)

	// ----- 数据类型
	fmt.Println()
	dataType()

	// ----- 局部变量
	fmt.Println()
	localVariable()

	// ----- 局部变量
	fmt.Println()
	globalVariable()

	// ----- 常量
	fmt.Println()
	constant()

	// ----- 枚举
	fmt.Println()
	constEnum()

	// ----- 枚举
	fmt.Println()
	iotaEnum()

	// ----- 运算符
	fmt.Println()
	d01.ArithmeticOperator()
	fmt.Println()
	d01.RelationalOperator()
	fmt.Println()
	d01.LogicalOperator()
	fmt.Println()
	d01.BitwiseOperator()
	fmt.Println()
	d01.AssignmentOperator()
	fmt.Println()
	d01.OtherOperators()

	// ----- 控制语句
	fmt.Println()
	d01.TestIf(13)
	fmt.Println()
	d01.TestSwitch("A")
	fmt.Println()
	d01.TestSelect()
	fmt.Println()
	d01.TestFor(0)

	// ----- 函数
	fmt.Println()
	d02.TestFunc()
	fmt.Println()
	d02.ValuePassed()
	fmt.Println()
	d02.ReferencePassed()
	fmt.Println()
	d02.TestClosure()
	fmt.Println()
	d02.TestMethod()

	// ----- 数组
	fmt.Println()
	d03.TestArray()
	// ----- 指针
	fmt.Println()
	d03.TestPointer()
	// ----- 结构
	fmt.Println()
	d03.TestStruct()
	// ----- 切片
	fmt.Println()
	d03.TestSlice()
	// ----- 循环
	fmt.Println()
	d03.TestRange()
	// ----- 集合
	fmt.Println()
	d03.TestMap()

	// ----- 递归
	fmt.Println()
	d04.TestRecursion()
	// ----- 接口
	fmt.Println()
	d04.TestInterface()
	// ----- 错误
	fmt.Println()
	d04.TestError()

	// ----- 纤程
	fmt.Println()
	d05.TestGoroutine()
	// ----- 同步安全sync
	fmt.Println()
	d05.TestSync()
	// ----- 原子操作
	fmt.Println()
	d05.TestAtomic()
	// ----- 通道
	fmt.Println()
	d05.TestChannel()
	// ----- 通道缓冲区
	fmt.Println()
	d05.TestChanBuffer()
	// ----- 通道关闭
	fmt.Println()
	d05.TestChanClose()

	fmt.Println()
	d05.TestSelect()
	fmt.Println()
	d05.TestTimeout()
}

/**
 * 将iota当作枚举使用
 */
const (
	// ----- 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1
	iotaA = iota
	iotaB = "测试跳过是否会+1"
	iotaC = iota
)

func iotaEnum() {
	fmt.Println("====================================== 枚举 ======================================")
	fmt.Println("iotaA:", iotaA, ", iotaB:", iotaB, "iotaC:", iotaC)
}

/**
 * 将常量当作枚举使用
 */
const (
	EAST  = "E"
	SOUTH = "S"
	WEST  = "W"
	NORTH = "N"
)

func constEnum() {
	fmt.Println("====================================== 枚举 ======================================")
	fmt.Println("东:", EAST, ", 南:", SOUTH, "西:", WEST, ", 北:", NORTH)
}

/**
 * 常量的声明与使用相关示例
 */
func constant() {
	fmt.Println("====================================== 常量 ======================================")
	// ----- 常量声明方式1：显示类型
	const constA1, constA2 string = "显示类型", "constA2"
	fmt.Println("constA1:", constA1, ", constA2:", constA2)
	// ----- 常量声明方式2：隐式类型
	const constB1, constB2 = "隐式类型", math.MaxInt32
	fmt.Println("constB1:", constB1, ", constB2:", constB2)
}

/**
 * 全局变量的声明与使用相关示例
 */
// ----- 全局变量的声明方式1：赋值操作
var globalA1, globalA2, globalA3 = "globalA1", "globalA2", "globalA3"

// ----- 全局变量的声明方式2：因式分解
var (
	globalB1 = "globalB1"
	globalB2 = "globalB2"
	globalB3 = "globalB3"
)

func globalVariable() {
	fmt.Println("==================================== 全局变量 =====================================")
	fmt.Println("globalA1:", globalA1, ", globalA2:", globalA2, ", globalA3:", globalA3)
	fmt.Println("globalB1:", globalB1, ", globalB2:", globalB2, ", globalB3:", globalB3)
}

/**
 * 局部变量的声明与使用相关示例
 */
func localVariable() {
	fmt.Println("==================================== 局部变量 =====================================")
	// ----- 变量声明方式1：一般声明
	var var1 string
	var1 = "指定变量类型，如果没有初始化，则变量默认为零值"
	fmt.Println("一般声明:", var1)

	// ----- 变量声明方式2：赋值操作
	var var2 = "根据值自行判定变量类型"
	fmt.Println("赋值操作:", var2)

	// ----- 变量声明方式3：简单声明
	var3 := "省略 var, 注意 := 左侧如果没有声明新的变量，就产生编译错误"
	fmt.Println("简单声明:", var3)

	// ----- 多变量声明方式1：一般声明
	var batchA1, batchA2, batchA3 string
	batchA1, batchA2, batchA3 = "batchA1", "batchA2", "batchA3"
	fmt.Println("batchA1:", batchA1, ", batchA2:", batchA2, ", batchA3:", batchA3)

	// ----- 多变量声明方式2：赋值操作
	var batchB1, batchB2, batchB3 = "batchB1", "batchB2", "batchB3"
	fmt.Println("batchB1:", batchB1, ", batchB2:", batchB2, ", batchB3:", batchB3)

	// ----- 多变量声明方式3：简单声明
	batchC1, batchC2, batchC3 := "batchC1", "batchC2", "batchC3"
	fmt.Println("batchC1:", batchC1, ", batchC2:", batchC2, ", batchC3:", batchC3)
}

/**
 * 数据类型相关示例
 */
func dataType() {
	fmt.Println("==================================== 数据类型 =====================================")
	fmt.Println("========== 布尔型 ==========")
	enabled := true
	disabled := 1 == 0
	fmt.Println("enabled: ", enabled, ", disabled: ", disabled)

	fmt.Println("========== 数值型 ==========")
	minUint8 := uint8(0)
	maxUint8 := math.MaxUint8
	fmt.Println("minUint8: ", minUint8, ", maxUint8: ", maxUint8)

	minUint16 := uint16(0)
	maxUint16 := math.MaxUint16
	fmt.Println("minUint16:", minUint16, ", maxUint16:", maxUint16)

	minUint32 := uint32(0)
	maxUint32 := math.MaxUint32
	fmt.Println("minUint32:", minUint32, ", maxUint32:", maxUint32)

	minUint64 := uint64(0)
	maxUint64 := uint64(math.MaxUint64)
	fmt.Println("minUint64:", minUint64, ", maxUint64:", maxUint64)

	minInt8 := math.MinInt8
	maxInt8 := math.MaxInt8
	fmt.Println("minInt8: ", minInt8, ", maxInt8: ", maxInt8)

	minInt16 := math.MinInt16
	maxInt16 := math.MaxInt16
	fmt.Println("minInt16:", minInt16, ", maxInt16:", maxInt16)

	minInt32 := math.MinInt32
	maxInt32 := math.MaxInt32
	fmt.Println("minInt32:", minInt32, ", maxInt32:", maxInt32)

	minInt64 := math.MinInt64
	maxInt64 := math.MaxInt64
	fmt.Println("minInt64:", minInt64, ", maxInt64:", maxInt64)

	fmt.Println("========== 浮点型 ==========")
	minFloat32 := float32(math.SmallestNonzeroFloat32)
	maxFloat32 := math.MaxFloat32
	fmt.Println("minFloat32:", minFloat32, ", maxFloat32:", maxFloat32)

	minFloat64 := float32(math.SmallestNonzeroFloat64)
	maxFloat64 := math.MaxFloat64
	fmt.Println("minFloat64:", minFloat64, ", maxFloat64:", maxFloat64)

	varComplex64 := complex64(0)
	varComplex128 := complex128(0)
	fmt.Println("varComplex64:", varComplex64, ", varComplex128:", varComplex128)

	fmt.Println("========== 其他型 ==========")
	vatByte := byte(0xFF)
	varRune := rune(math.MaxInt32)
	varUint := uint(math.MaxUint64)
	varInt := math.MaxInt64
	fmt.Println("vatByte:", vatByte, ", varRune:", varRune, ", varUint:", varUint, ", varInt:", varInt)
}
