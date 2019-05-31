package service

import (
	"fmt"
	"github.com/kataras/iris/core/errors"
	"goIM/model"
	"goIM/util"
	"log"
	"math/rand"
	"time"
)

type UserService struct {
}

// 注册函数
func (s *UserService) Register(
	mobile, // 手机
	plainpwd, // 明文密码
	nickname, // 昵称
	avater, sex string) (user model.User, err error) {

	//检测手机号码是否存在,
	tmp := model.User{}
	_, err = DBEngin.Where("mobile=? ", mobile).Get(&tmp)
	if err != nil {
		return tmp, err
	}

	if tmp.Id > 0 {
		return tmp, errors.New("该手机号已经注册")
	}
	//否则拼接插入数据
	tmp.Mobile = mobile
	tmp.Avatar = avater
	tmp.Nickname = nickname
	tmp.Sex = sex
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	//md5 加密
	tmp.Passwd = util.MakePasswd(plainpwd, tmp.Salt)
	tmp.Createat = time.Now()
	// token 可以是一个随机数
	tmp.Token = fmt.Sprintf("%08d", rand.Int31())

	//插入 InserOne
	_, err = DBEngin.InsertOne(&tmp)

	return tmp, err
}

func (s *UserService) Login(
	mobile,
	plainwd string) (user model.User, err error) {
	//首先通过手机号查询用户
	tmp := model.User{}
	DBEngin.Where("mobile=? ", mobile).Get(&tmp)
	// 如果没有找到
	if tmp.Id == 0 {
		return tmp, errors.New("该用户不存在")
	}
	// 查询完成比对密码
	if !util.ValidatePasswd(plainwd, tmp.Salt, tmp.Passwd) {
		return tmp, errors.New("密码不正确")
	}
	//比对密码是否正确
	str := fmt.Sprintf("%d", time.Now().Unix())
	token := util.MD5Encode(str)
	//刷新token，保证安全性
	tmp.Token = token
	// 返回数据
	_, err = DBEngin.ID(tmp.Id).Cols("token").Update(&tmp)
	if err != nil {
		log.Println("[FAIL] DBEngin.ID(tmp.Id).Cols(\"token\").Update(&tmp)")
	}

	return tmp, nil
}

//查找某个用户
func (s *UserService) Find(
	userId int64) (user model.User) {

	//首先通过手机号查询用户
	tmp := model.User{}
	DBEngin.ID(userId).Get(&tmp)
	return tmp
}
