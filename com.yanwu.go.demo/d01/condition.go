package d01

import "fmt"

/**
 * @Author Baofeng Xu
 * @Date 2021/5/13 9:56
 *
 * Description: 条件控制语句
 **/

func TestSelect() {
	fmt.Println("====================================== switch ======================================")
	var chanA, chanB, chanC chan int
	var intA, intB int
	select {
	// ----- 从通道chanA读取数据赋值给intA
	case intA = <-chanA:
		fmt.Println("chanA reader intA:", intA)
	// ----- 将intB发送至通道chanB
	case chanB <- intB:
		fmt.Println("chanB sender intB:", intB)
	// ----- 从通道chanC中读取数据赋值给intC和ok
	case intC, ok := <-chanC:
		if ok {
			fmt.Println("chanC reader intC:", intC)
		} else {
			fmt.Println("chanC closed.")
		}
	default:
		fmt.Println("no communication.")
	}
}

func TestSwitch(direction string) {
	fmt.Println("====================================== switch ======================================")
	switch direction {
	case "E":
		fmt.Println("朝向为 > EAST:", direction)
	case "S":
		fmt.Println("朝向为 > SOUTH:", direction)
	case "W":
		fmt.Println("朝向为 > WEST:", direction)
	case "N":
		fmt.Println("朝向为 > NORTH:", direction)
	default:
		fmt.Println("方向错误:", direction)
	}
}

func TestIf(i int) {
	fmt.Println("====================================== if ==========================================")
	if i%2 == 0 {
		fmt.Println("i 是偶数:", i)
	} else {
		fmt.Println("i 是偶数:", i)
	}
}
