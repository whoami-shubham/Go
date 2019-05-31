package main

import (
	"fmt"
	"os"
	"text/template"
)
var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

/*
  init() will parse templates  once .
  template.Must() returns *template.Template and check for error
*/

func main(){

	 err := tpl.ExecuteTemplate(os.Stdout,"index.gohtml",nil)
	 if err != nil {
		 fmt.Println(err)
	 }
}