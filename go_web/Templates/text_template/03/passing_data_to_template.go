package main

import (
	"os"
	"fmt"
	"text/template"
)

var tpl *template.Template

func init(){
   tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main(){

	err := tpl.Execute(os.Stdout,"title")

	if err != nil {
		fmt.Println(err)
	}
}