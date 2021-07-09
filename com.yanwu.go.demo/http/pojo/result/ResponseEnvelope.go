package result

import (
	"encoding/json"
	"log"
	"net/http"
)

/**
 * @Author Baofeng Xu
 * @Date 2021/7/9 10:55
 *
 * Description: HTTP接口结果响应
 **/

type ResponseEnvelope struct {
	Code    int         `json:"code"`    // 接口响应码
	Message string      `json:"message"` // 接口错误提示语
	Data    interface{} `json:"data"`    // 数据
}

// Success 接口调用成功
func Success(data interface{}) *ResponseEnvelope {
	return getInstance(200, data, "")
}

// Failed 接口调用失败
func Failed(message string) *ResponseEnvelope {
	return getInstance(500, nil, message)
}

// Result 接口调用失败
func Result(w http.ResponseWriter, result *ResponseEnvelope) {
	bytes, _ := json.Marshal(result)
	write, err := w.Write(bytes)
	if err != nil {
		log.Println("Response result error.", write, err)
	}
}

func getInstance(code int, data interface{}, message string) *ResponseEnvelope {
	return &ResponseEnvelope{Code: code, Data: data, Message: message}
}
