package d03

import (
	constant "../constant"
	"fmt"
)

/**
 * @Author Baofeng Xu
 * @Date 2021/5/13 14:42
 *
 * Description: 结构
 **/

// User 用户
type User struct {
	username string  // 用户名
	password string  // 密码
	sex      bool    // 性别 [true: 男; false: 女]
	age      uint8   // 年龄
	height   float32 // 身高
	weight   float32 // 体重
}

func TestStruct() {
	fmt.Println("====================================== 结构 ======================================")
	var yanWu User
	yanWu.username = "YanWu"
	yanWu.password = "12138"
	yanWu.sex = constant.MAN
	yanWu.age = 29
	yanWu.height = 1.72
	yanWu.weight = 71.2
	fmt.Println("yanWu:", yanWu)
	printUser(yanWu)

	lotus := User{"lotus", "12138", constant.WOMAN, 29, 1.68, 51}
	printUser(lotus)

	wenXin := User{username: "WenXin", sex: constant.MAN, age: 3}
	printUser(wenXin)

	wenFu := User{username: "WenFu", sex: constant.WOMAN, age: 2}
	printUser(wenFu)
}

func printUser(user User) {
	fmt.Println("username:", user.username, ", password:", user.password, ", sex:", user.sex, ", age:", user.age, ", height:", user.height, ", weight:", user.weight)
}
