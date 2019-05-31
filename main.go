package main

import (
	"fmt"
	"goIM/ctrl"
	"html/template"
	"log"
	"net/http"
)

func RegisterView() {
	tpl, err := template.ParseGlob("view/**/*")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, v := range tpl.Templates() {
		tplname := v.Name()

		http.HandleFunc(tplname,
			func(writer http.ResponseWriter,
				request *http.Request) {
				tpl.ExecuteTemplate(writer, tplname, nil)
			})
	}
	fmt.Println("Templates Rendering Succ!")
}

func main() {
	// 绑定请求和处理函数
	http.HandleFunc("/user/login", ctrl.UserLogin)
	http.HandleFunc("/user/register", ctrl.UserRegister)
	http.HandleFunc("/contact/loadcommunity", ctrl.LoadCommunity)
	http.HandleFunc("/contact/loadfriend", ctrl.LoadFriend)
	http.HandleFunc("/contact/joincommunity", ctrl.JoinCommunity)
	http.HandleFunc("/contact/addfriend", ctrl.Addfriend)
	http.HandleFunc("/chat", ctrl.Chat)
	http.HandleFunc("/attach/upload", ctrl.Upload)

	// 提供静态资源使用
	// http.Handle("/", http.FileServer(http.Dir("."))) // 全局设置，maingo可以访问
	// 提供指定目录的静态文件支持
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	http.Handle("/mnt/", http.FileServer(http.Dir(".")))

	// 查找并渲染模板
	RegisterView()

	// 启动web服务器
	http.ListenAndServe(":8080", nil)

	select {}
}
