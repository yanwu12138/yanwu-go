package user

import (
	constant "../../../constant"
	model "../../data/model"
	result "../../pojo/result"
	userService "../../service/user"
	util "../../util"
	"log"
	"net/http"
	"strconv"
	"time"
)

/**
 * @Author Baofeng Xu
 * @Date 2021/7/9 15:22
 *
 * Description: 用户相关接口
 **/

// Create 创建用户
func Create(w http.ResponseWriter, req *http.Request) {
	// ===== 处理参数
	err := req.ParseForm()
	if err != nil {
		result.Result(w, result.Failed("参数获取失败"))
		return
	}
	// ----- 用户名
	name, found1 := req.Form["name"]
	if !found1 {
		log.Println("create user error, because username is empty.")
		return
	}
	user := model.YanwuUser{Name: name[0], CreatedAt: time.Now(), UpdatedAt: time.Now()}
	// ----- 邮箱
	emile, _ := req.Form["emile"]
	if len(emile) != 0 {
		user.Email = emile[0]
	}
	// ----- 年龄
	ageStr, _ := req.Form["age"]
	if len(ageStr) != 0 {
		age, _ := strconv.ParseUint(ageStr[0], 10, 64)
		user.Age = age
	}
	// ----- 密码
	password, _ := req.Form["password"]
	if len(password) == 0 {
		user.Password = util.Encrypt(constant.PASSWORD)
	} else {
		user.Password = util.Encrypt(password[0])
	}
	log.Println("create user:", user)
	// ----- 执行登录操作
	user.ID = userService.Create(&user)
	if user.ID == 0 {
		result.Result(w, result.Failed("服务器内部错误"))
	} else {
		result.Result(w, result.Success(user))
	}
}

// Update 更新用户信息
func Update(w http.ResponseWriter, req *http.Request) {

}

// GetById 根据用户ID查看用户详情
func GetById(w http.ResponseWriter, req *http.Request) {

}

// All 查看所有用户信息
func All(w http.ResponseWriter, req *http.Request) {

}

// Delete 删除用户
func Delete(w http.ResponseWriter, req *http.Request) {

}

// UpdatePassword 删除用户
func UpdatePassword(w http.ResponseWriter, req *http.Request) {

}
