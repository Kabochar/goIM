package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type H struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func RespFail(w http.ResponseWriter, msg string) {
	Resp(w, -1, nil, msg)
}

func RespOk(w http.ResponseWriter, data interface{}, msg string) {
	Resp(w, 0, data, msg)
}

func Resp(writer http.ResponseWriter, code int, data interface{}, msg string) {
	//设置header 为幻S0N 默认的text/html，所以特别指出返回的
	//为application/json
	//设置header
	writer.Header().Set("Content-Type", "application/json")
	//设置200状态码
	writer.WriteHeader(http.StatusOK)
	//输出
	//定义一个结构体
	h := H{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	writer.Write(ret)
}
