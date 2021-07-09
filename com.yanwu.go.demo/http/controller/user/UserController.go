package user

import (
	constant "../../../constant"
	model "../../data/model"
	result "../../pojo/result"
	userService "../../service/user"
	util "../../util"
	"log"
	"net/http"
	"time"
)

/**
 * @Author Baofeng Xu
 * @Date 2021/7/9 15:22
 *
 * Description: 用户相关接口
 **/

// Create 创建用户
func Create(response http.ResponseWriter, request *http.Request) {
	requestFrom := util.RequestFrom(response, request)
	if !requestFrom {
		return
	}
	// ----- 用户名
	name := util.RequestParamToString("name", false, response, request)
	user := model.YanwuUser{Name: name, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	// ----- 邮箱
	emile := util.RequestParamToString("emile", true, response, request)
	user.Email = emile
	// ----- 年龄
	age := util.RequestParamToUint("age", true, response, request)
	user.Age = uint64(age)
	// ----- 密码
	password := util.RequestParamToString("password", true, response, request)
	if len(password) == 0 {
		user.Password = util.Encrypt(constant.PASSWORD)
	} else {
		user.Password = util.Encrypt(password)
	}
	log.Println("create user:", user)
	// ----- 执行登录操作
	user.ID = userService.Create(&user)
	if user.ID == 0 {
		result.Result(response, result.Failed("服务器内部错误"))
	} else {
		result.Result(response, result.Success(user))
	}
}

// Update 更新用户信息
func Update(response http.ResponseWriter, request *http.Request) {
	requestFrom := util.RequestFrom(response, request)
	if !requestFrom {
		return
	}
}

// GetById 根据用户ID查看用户详情
func GetById(response http.ResponseWriter, request *http.Request) {
	// ===== 处理参数
	requestFrom := util.RequestFrom(response, request)
	if !requestFrom {
		return
	}
	// ----- 参数校验
	id := util.RequestParamToUint("id", false, response, request)
	user := userService.GetById(id)
	result.Result(response, result.Success(user))
	log.Println("get user by id, id:", id, " user:", user)
}

// All 查看所有用户信息
func All(response http.ResponseWriter, request *http.Request) {
	result.Result(response, result.Success(userService.AllUser()))
}

// Delete 删除用户
func Delete(response http.ResponseWriter, request *http.Request) {

}

// UpdatePassword 删除用户
func UpdatePassword(response http.ResponseWriter, request *http.Request) {

}
