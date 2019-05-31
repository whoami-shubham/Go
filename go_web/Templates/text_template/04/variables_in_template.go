package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

func init(){
  tpl = template.Must(template.ParseFiles("one.gohtml","two.gohtml"))
}

func main(){
	data := []string{"one","two","three","four","five"}
	struct_data := struct{
		Name string
		Age  string
	}{
		"john Doe",
		"20",
	}
	err := tpl.ExecuteTemplate(os.Stdout,"one.gohtml",data)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("\n-------------------2nd Template ------------------")
	err = tpl.ExecuteTemplate(os.Stdout,"two.gohtml",struct_data)
	if err != nil{
		fmt.Println(err)
	}
}