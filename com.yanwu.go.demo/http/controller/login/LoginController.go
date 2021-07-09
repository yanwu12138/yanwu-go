package login

import (
	pojo "../../pojo"
	result "../../pojo/result"
	util "../../util"
	"encoding/json"
	"log"
	"net/http"
)

/**
 * @Author Baofeng Xu
 * @Date 2021/7/9 10:49
 *
 * Description: 用户登录等出接口
 **/

// Login 用户登录
func Login(response http.ResponseWriter, request *http.Request) {
	// ----- 处理参数
	requestFrom := util.RequestFrom(response, request)
	if !requestFrom {
		return
	}

	loginUser := pojo.LoginUser{
		Account:  util.RequestParamToString("account", false, response, request),
		Password: util.RequestParamToString("password", false, response, request),
	}

	log.Println("user Login, param:", loginUser)
	result.Result(response, result.Success(loginUser))
}

// Logout 用户登出
func Logout(w http.ResponseWriter, request *http.Request) {
	log.Println("user Logout, param:", request)
	bytes, _ := json.Marshal(result.Success("Login"))
	w.Write(bytes)
}
