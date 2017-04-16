package main

import (
	//"fmt"
	"net/http"

	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
	"github.com/xiaozi0lei/WanYingNote/controller"
	"log"
)

func main() {

	// 返回 http.Handler
	router := mux.NewRouter()

	// 注册各个 handler，路由的功能
	router.HandleFunc("/", controller.Index)
	router.HandleFunc("/name", controller.GetName)
	//router.HandleFunc("/", HomeHandler)

	n := negroni.Classic()
	// 将 http.Handler 传入中间件栈，handler 按加入顺序先后执行
	n.UseHandler(router)

	//
	log.Fatal(http.ListenAndServe(":3000", n))
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	controller.Index(rw, r)
	//vars := mux.Vars(r)
	//w.WriteHeader(http.StatusOK)
	//fmt.Fprintf(w, "%v", "Welcome to my home page!")
	//
	//fmt.Fprintf(w, "Category: %v\n", vars)
	////fmt.Fprintf(w, "Category: %v\n", vars["category"])
}