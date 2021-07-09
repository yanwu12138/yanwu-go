package d03

import "fmt"

/**
 * @Author Baofeng Xu
 * @Date 2021/5/13 16:18
 *
 * Description: 集合
 **/

func TestMap() {
	fmt.Println("====================================== 集合 ======================================")
	var mapA map[string]int
	mapA = make(map[string]int)
	mapA["yanWu"] = 29
	mapA["lotus"] = 29
	fmt.Println("mapA:", mapA, ", length:", len(mapA))

	age, ok := mapA["yanWu"]
	if ok {
		fmt.Println("yanWu age:", age)
	} else {
		fmt.Println("yanWu 不存在")
	}

	fmt.Println("----------")
	mapB := map[string]string{"湖北": "武汉", "湖南": "长沙", "浙江": "杭州"}
	for key, value := range mapB {
		fmt.Println(key, "的省会:", value)
	}

	delete(mapB, "浙江")
	fmt.Println("----------")
	for key, value := range mapB {
		fmt.Println(key, "的省会:", value)
	}

	mapB["浙江"] = "杭州"
	fmt.Println("----------")
	for key, value := range mapB {
		fmt.Println(key, "的省会:", value)
	}
}
