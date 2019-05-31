package main

import (
	"os"
	"fmt"
	"text/template"
)

func main(){

	tpl,err := template.ParseFiles("index.gohtml")
	/*
	  we can also use ParseGlob to parse files

	*/
	if err != nil {
		fmt.Println(err)
	}

	err = tpl.Execute(os.Stdout,nil)
	
	/*
	  we can also use ExecuteTemplate()
	  to execute specific template 

	*/

	if err != nil {
		fmt.Println(err)
	}

}