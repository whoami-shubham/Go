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

type handler int

func (h handler) ServeHTTP(w http.ResponseWriter,req *http.Request){
	if req.Method == "GET" {
		err := tpl.Execute(w,nil)
		if err != nil{
			fmt.Fprintln(w,err)
		}
	}else{
		req.ParseForm()
		err := tpl.Execute(w,req.Form)
        if err != nil{
			fmt.Fprintln(w,err)
		}
	}
}

func main(){
	
	var h handler 

	http.ListenAndServe(":8080",h)


}