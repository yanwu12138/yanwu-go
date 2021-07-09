package d04

import "fmt"

/**
 * @Author Baofeng Xu
 * @Date 2021/5/13 16:31
 *
 * Description: 递归
 **/

func TestRecursion() {
	fmt.Println("====================================== 递归 ======================================")
	fmt.Println("5 的阶乘:", factorial(5))
}

func factorial(num uint64) uint64 {
	if num > 0 {
		return num * factorial(num-1)
	}
	return 1
}
