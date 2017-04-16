package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/{name}", HomeHandler)

	n := negroni.Classic()
	n.UseHandler(router)

	http.ListenAndServe(":3000", n)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to my home page!")
	
	fmt.Fprintf(w, "Category: %v\n", vars)
	//fmt.Fprintf(w, "Category: %v\n", vars["category"])
}