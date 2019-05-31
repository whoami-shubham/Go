package main

import (
	"os"
	"fmt"
	"text/template"
)

var tpl *template.Template
var fm = template.FuncMap{
  "sum": Sum,
   "mult": Mult,
}

func init(){
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
}

func main(){
	
	data := []int{1,2,3,4,5,6,7,8}
	err := tpl.Execute(os.Stdout,data)

	if err!=nil {
		fmt.Println(err)
	}

}

func Sum(n []int)(int){
	 ans := 0
	 for _,value := range n {
		 ans = ans + value
	 }
	 return ans
}

func Mult(n []int)(int){
	ans := 1
	for _,value := range n{
		ans = ans*value
	}
	return ans
}