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

	user, err := userService.Login(mobile, passwd)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOk(writer, user, "")
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
