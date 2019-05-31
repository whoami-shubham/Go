package main

import (

	"fmt"
	"net/http"
	"html/template"

)

var tpl *template.Template

func init(){
    tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main(){

	http.HandleFunc("/",home)
	http.HandleFunc("/about/",about)
	http.HandleFunc("/contact/",contact)

	http.ListenAndServe(":8080",nil)  

}

func home(w http.ResponseWriter, req *http.Request){
	fmt.Println(req.Method,req.URL, " home()")

	err := tpl.ExecuteTemplate(w,"index.gohtml",req.URL)

	if err != nil {
		fmt.Fprintln(w,err)
	}

}

func about(w http.ResponseWriter, req *http.Request){
	fmt.Println(req.Method,req.URL, " about()")
	err := tpl.ExecuteTemplate(w,"index.gohtml",req.URL)

	if err != nil {
		fmt.Fprintln(w,err)
	}
	
}

func contact(w http.ResponseWriter, req *http.Request){
	fmt.Println(req.Method,req.URL, " contact()")
	err := tpl.ExecuteTemplate(w,"index.gohtml",req.URL)

	if err != nil {
		fmt.Fprintln(w,err)
	}
	
}