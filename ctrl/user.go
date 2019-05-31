package ctrl

import (
	"fmt"
	"goIM/model"
	"goIM/service"
	"goIM/util"
	"math/rand"
	"net/http"
)

func UserLogin(writer http.ResponseWriter, request *http.Request) {
	// 1 获取前端传递的参数
	// 解析参数
	// 如何获得参数
	request.ParseForm()

	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")

	loginok := false
	if mobile == "12610517222" && passwd == "123456" {
		loginok = true
	}
	if loginok {
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "test"
		util.RespOk(writer, data, "")
	} else {
		util.RespFail(writer, "passwd error!")
	}
}

var userService service.UserService

func UserRegister(writer http.ResponseWriter,
	request *http.Request) {

	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	plainpwd := request.PostForm.Get("passwd")
	nickname := fmt.Sprintf("user%06d", rand.Int31())
	avatar := ""
	sex := model.SEX_UNKNOW

	user, err := userService.Register(mobile, plainpwd, nickname, avatar, sex)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOk(writer, user, "")
	}
}
