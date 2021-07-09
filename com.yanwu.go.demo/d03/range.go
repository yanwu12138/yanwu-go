package d03

import "fmt"

/**
 * @Author Baofeng Xu
 * @Date 2021/5/13 16:04
 *
 * Description: 范围
 **/

func TestRange() {
	fmt.Println("====================================== 范围 ======================================")
	// ----- 迭代数组
	sliceA := []int{1, 2, 3, 4, 5}
	for index, item := range sliceA {
		fmt.Println("slice > index:", index, ", item", item)
	}

	// ----- 迭代map
	mapA := map[string]int{"yanWu": 29, "lotus": 29, "wenXin": 3, "wenFu": 2}
	for key, value := range mapA {
		fmt.Println("map > key:", key, ", value", value)
	}
}
