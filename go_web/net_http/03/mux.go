package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type handler int

var tpl *template.Template

func init(){

	tpl = template.Must(template.ParseFiles("index.gohtml"))
}


func main(){

	var home handler
	
	mux := http.NewServeMux()  // it returns *ServeMux() that have functions Handle and HandleFunc
	mux.Handle("/",home)
	mux.Handle("/about/",http.HandlerFunc(about)) // handle /about/*
	mux.HandleFunc("/contact",contact)            // only handle /contact
	

	http.ListenAndServe(":8080",mux)

}

func (h handler) ServeHTTP(w http.ResponseWriter,req *http.Request){
	fmt.Println(req.Method, req.URL,"  inside home() ")
	err := tpl.Execute(w,req.URL)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w,err)
	}
}
func about(w http.ResponseWriter,req *http.Request){
	fmt.Println(req.Method, req.URL,"  inside about() ")
	err := tpl.Execute(w,req.URL)
    
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w,err)
	}
}

func contact(w http.ResponseWriter,req *http.Request){
	fmt.Println(req.Method, req.URL,"  inside contact() ")
	err := tpl.Execute(w,req.URL)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w,err)
	}
}

