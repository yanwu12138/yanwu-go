package main

import (
	login "./controller/login"
	user "./controller/user"
	"log"
	"net/http"
)

/**
 * @Author Baofeng Xu
 * @Date 2021/7/9 11:09
 *
 * Description: HTTP服务入口
 **/

func init() {
	log.SetPrefix("yanwu go demo for HTTP: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	log.Println("This is webserver base!")

	RequestAgent()

	//服务器要监听的主机地址和端口号
	err := http.ListenAndServe("127.0.0.1:8081", nil)
	if err != nil {
		log.Println("ListenAndServe error: ", err.Error())
	}
}

func RequestAgent() {
	// ---------- 用户登录登出 ---------- //
	http.HandleFunc("/login", login.Login)
	http.HandleFunc("/logout", login.Logout)
	// ---------- 用户操作 ---------- //
	http.HandleFunc("/user/create", user.Create)
	http.HandleFunc("/user/update", user.Update)
	http.HandleFunc("/user/getById", user.GetById)
	http.HandleFunc("/user/all", user.All)
	http.HandleFunc("/user/delete", user.Delete)
	http.HandleFunc("/user/updatePassword", user.UpdatePassword)
}
