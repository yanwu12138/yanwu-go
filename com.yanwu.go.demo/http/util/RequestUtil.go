package util

import (
	result "../pojo/result"
	"log"
	"net/http"
	"strconv"
)

/**
 * @Author Baofeng Xu
 * @Date 2021/7/9 17:47
 *
 * Description: HTTP请求时的相关工具
 **/

// RequestFrom 处理request，方便后续从request中获取参数
func RequestFrom(response http.ResponseWriter, request *http.Request) bool {
	err := request.ParseForm()
	if err == nil {
		return true
	}
	result.Result(response, result.Failed("参数获取失败"))
	return false
}

// RequestParamToString 从request中获取string类型的参数
func RequestParamToString(name string, allowedNil bool, response http.ResponseWriter, request *http.Request) string {
	return requestParam(name, allowedNil, response, request)[0]
}

// RequestParamToInt 从request中获取int类型参数
func RequestParamToInt(name string, allowedNil bool, response http.ResponseWriter, request *http.Request) int {
	param := requestParam(name, allowedNil, response, request)[0]
	paramInt, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		log.Println("参数:", name, "格式不正确")
		result.Result(response, result.Failed("参数: "+name+" 格式不正确"))
		panic("参数: " + name + " 格式不正确")
	}
	return int(paramInt)
}

// RequestParamToUint 从request中获取uint类型参数
func RequestParamToUint(name string, allowedNil bool, response http.ResponseWriter, request *http.Request) uint {
	param := requestParam(name, allowedNil, response, request)[0]
	paramUint, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		log.Println("参数:", name, "格式不正确")
		result.Result(response, result.Failed("参数: "+name+" 格式不正确"))
		panic("参数: " + name + " 格式不正确")
	}
	return uint(paramUint)
}

// RequestParamToBool 从request中获取bool类型的参数
func RequestParamToBool(name string, allowedNil bool, response http.ResponseWriter, request *http.Request) bool {
	param := requestParam(name, allowedNil, response, request)[0]
	paramBool, err := strconv.ParseBool(param)
	if err != nil {
		log.Println("参数:", name, "格式不正确")
		result.Result(response, result.Failed("参数: "+name+" 格式不正确"))
		panic("参数: " + name + " 格式不正确")
	}
	return paramBool
}

// RequestParamToFloat 从request中获取float类型的参数
func RequestParamToFloat(name string, allowedNil bool, response http.ResponseWriter, request *http.Request) float64 {
	param := requestParam(name, allowedNil, response, request)[0]
	paramFloat, err := strconv.ParseFloat(param, 64)
	if err != nil {
		log.Println("参数:", name, "格式不正确")
		result.Result(response, result.Failed("参数: "+name+" 格式不正确"))
		panic("参数: " + name + " 格式不正确")
	}
	return paramFloat
}

/***
 * requestParam 	从request中获取参数
 * name 			属性名
 * allowedNil 		是否允许为nil: [true: 允许; false: 不允许]
 */
func requestParam(name string, allowedNil bool, response http.ResponseWriter, request *http.Request) []string {
	param, found := request.Form[name]
	if (!found || len(param) == 0) && !allowedNil {
		log.Println("参数:", name, "不能为空")
		result.Result(response, result.Failed("参数: "+name+" 不能为空"))
		panic("参数: " + name + " 不能为空")
	}
	return param
}
