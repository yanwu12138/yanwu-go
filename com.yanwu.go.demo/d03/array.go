package d03

import "fmt"

/**
 * @Author Baofeng Xu
 * @Date 2021/5/13 11:45
 *
 * Description: 数组
 **/

func TestArray() {
	fmt.Println("====================================== 数组 ======================================")
	var arrayA [10]int
	fmt.Println("arrayA:", arrayA, ", length:", len(arrayA))
	arrayA[0] = 10
	arrayA[7] = 10
	fmt.Println("arrayA:", arrayA, ", length:", len(arrayA))

	var arrayB = [...]string{"YanWu", "lotus", "WenXin", "WenFu"}
	fmt.Println("arrayB:", arrayB, ", length:", len(arrayB))

	arrayC := [...]string{"YanWu", "lotus", "WenXin", "WenFu"}
	fmt.Println("arrayC:", arrayC, ", length:", len(arrayC))

	for i := 0; i < len(arrayC);i++ {
		fmt.Println("arrayC > index:", i, ", value:", arrayC[i])
	}
}
