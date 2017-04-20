package main

import (
	"net/http"
	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
	"github.com/xiaozi0lei/WanYingNote/controller"
	"log"
)

func main() {

	// ------ 1. 创建各种路由 ------ //
	// mux 是一个 router 的中间件，返回 http.Handler，
	// 比原生的 handler 有很多更方便的用法，支持丰富的匹配请求的方式
	// 支持正则匹配
	// 支持嵌套路由
	router := mux.NewRouter()
	// 注册各个 handler，路由的功能
	router.HandleFunc("/", controller.Index)
	router.HandleFunc("/name", controller.GetName)
	//router.HandleFunc("/", HomeHandler)

	// ------ 2. 接入各种中间件 ------ //
	// 默认支持 Recovery、Logger、Static 中间件
	n := negroni.New()
	// 设置 public 的映射路径及路由前缀
	publicContent := negroni.NewStatic(http.Dir("./public"))
	publicContent.Prefix = "/public"
	n.Use(publicContent)
	// 设置 public 的映射路径及路由前缀
	favicon := negroni.NewStatic(http.Dir("./favicon.ico"))
	favicon.Prefix = "/favicon.ico"
	n.Use(favicon)
	// 使用 negroni 自带的 Logger
	n.Use(negroni.NewLogger())
	// 使用 negroni 自带的 Recovery
	n.Use(negroni.NewRecovery())
	// 将 http.Handler 传入中间件栈，handler 按加入顺序先后执行
	n.UseHandler(router)

	// ------ 3. 加载所有路由和中间件，启动服务器 ------ //
	// 启动服务器，监听 3000 端口
	log.Println("Starting the server ...")
	log.Fatal(http.ListenAndServe(":60006", n))

}