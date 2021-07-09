package d03

import (
	"fmt"
)

/**
 * @Author Baofeng Xu
 * @Date 2021/5/13 15:11
 *
 * Description: 切片
 **/

func TestSlice() {
	fmt.Println("====================================== 切片 ======================================")
	var sliceA = make([]int, 2, 10)
	fmt.Println("sliceA:", sliceA, ", length:", len(sliceA), ", capacity:", cap(sliceA))
	sliceA[0] = 10
	fmt.Println("sliceA:", sliceA, ", length:", len(sliceA), ", capacity:", cap(sliceA))

	arr := [...]int{1, 2, 3, 4, 5}

	sliceB := arr[:]
	fmt.Println("sliceB:", sliceB, ", length:", len(sliceB), ", capacity:", cap(sliceB))

	sliceC := arr[0:]
	fmt.Println("sliceC:", sliceC, ", length:", len(sliceC), ", capacity:", cap(sliceC))

	sliceD := arr[:5]
	fmt.Println("sliceD:", sliceD, ", length:", len(sliceD), ", capacity:", cap(sliceD))

	sliceE := arr[0:5]
	fmt.Println("sliceE:", sliceE, ", length:", len(sliceE), ", capacity:", cap(sliceE))

	sliceF := arr[1:4]
	fmt.Println("sliceF:", sliceF, ", length:", len(sliceF), ", capacity:", cap(sliceF))

	sliceG := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("sliceG:", sliceG, ", length:", len(sliceG), ", capacity:", cap(sliceG))
	fmt.Println("sliceG[1:4]:", sliceG[1:4])
	fmt.Println("sliceG[:3]:", sliceG[:3])
	fmt.Println("sliceG[5:]:", sliceG[5:])
	// ----- append(): 追加元素
	sliceG = append(sliceG, 10)
	fmt.Println("sliceG:", sliceG, ", length:", len(sliceG), ", capacity:", cap(sliceG))
	sliceG = append(sliceG, 11, 12)
	fmt.Println("sliceG:", sliceG, ", length:", len(sliceG), ", capacity:", cap(sliceG))
	// ----- copy(): 拷贝切片
	sliceH := make([]int, len(sliceG), cap(sliceG))
	copy(sliceH, sliceG)
	fmt.Println("sliceH:", sliceH, ", length:", len(sliceH), ", capacity:", cap(sliceH))
}
