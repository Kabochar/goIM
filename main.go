package main

import (
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
}

func main() {
	// 绑定请求和处理函数
	http.HandleFunc("/user/login", ctrl.UserLogin)
	http.HandleFunc("/user/register", ctrl.UserRegister)

	// 提供静态资源使用
	// http.Handle("/", http.FileServer(http.Dir("."))) // 全局设置，maingo可以访问
	// 提供指定目录的静态文件支持
	http.Handle("/asset/", http.FileServer(http.Dir(".")))

	// 查找并渲染模板
	RegisterView()

	// 启动web服务器
	http.ListenAndServe(":8080", nil)

	select {}
}
