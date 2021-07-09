package d04

import "fmt"

/**
 * @Author Baofeng Xu
 * @Date 2021/5/13 17:13
 *
 * Description: 错误
 **/

func TestError() {
	fmt.Println("====================================== 错误 ======================================")
	result, err := div(10, 0)
	if err != "" {
		fmt.Println("err:", err)
	} else {
		fmt.Println("10/0 result:", result)
	}
}

func div(intA, intB int) (result int, err string) {
	if intB == 0 {
		runtimeErr := RuntimeError{100, "除数不能为0"}
		err = runtimeErr.Error()
		return
	} else {
		return intA / intB, ""
	}
}

type RuntimeError struct {
	code    uint   // 错误码
	message string // 错误信息
}

func (e RuntimeError) Error() string {
	return fmt.Sprintf("error >> code: %v, message: %v", e.code, e.message)
}
