package d01

import "fmt"

/**
 * @Author Baofeng Xu
 * @Date 2021/5/13 9:56
 *
 * Description: 循环控制语句
 **/

// TestFor for循环
func TestFor(intA int) {
	fmt.Println("====================================== for ======================================")
	for intA < 20 {
		if intA > 15 {
			// ----- 跳出循环
			break
		}
		if intA == 10 {
			// ----- 跳转到指定的行: AAA, 从AAA开始往下执行
			goto AAA
		}
		if intA%2 == 0 {
			intA++
			// ----- 跳过本次循环
			continue
		} else {
			fmt.Println("intA:", intA)
		}
		intA++
	}

AAA:
	fmt.Println("goto AAA, intA:", intA)
}
