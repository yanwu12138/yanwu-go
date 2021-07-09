package login

import (
	pojo "../../pojo"
	result "../../pojo/result"
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
func Login(w http.ResponseWriter, req *http.Request) {
	// ----- 处理参数
	err := req.ParseForm()
	if err != nil {
		result.Result(w, result.Failed("参数获取失败"))
		return
	}
	account, found1 := req.Form["account"]
	password, found2 := req.Form["password"]
	if !(found1 && found2) {
		log.Println("login error. param > account:", account, "password:", password)
		result.Result(w, result.Failed("参数错误"))
		return
	}
	loginUser := pojo.LoginUser{Account: account[0], Password: password[0]}
	log.Println("user Login, param:", loginUser)
	result.Result(w, result.Success(loginUser))
}

// Logout 用户登出
func Logout(w http.ResponseWriter, req *http.Request) {
	log.Println("user Logout, param:", req)
	bytes, _ := json.Marshal(result.Success("Login"))
	w.Write(bytes)
}
