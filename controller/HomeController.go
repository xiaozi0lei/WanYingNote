package controller

import(
	"fmt"
	"net/http"
	"html/template"
	"log"
)

func Index(rw http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./view/home/index.html"))
	err := tpl.Execute(rw, nil)

	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

	fmt.Println("hello, world!")
}

func GetName(rw http.ResponseWriter, r *http.Request) {

	rw.Write([]byte("GuoLei"))
	fmt.Println(r.Cookies())
	fmt.Println("GuoLei!")
}