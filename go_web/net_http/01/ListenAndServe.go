package main

import (
	"fmt"
	"net/http"
)

type handler int
func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request){
	fmt.Fprintln(w," Hello World !")
} 
func main(){
	var h handler
	http.ListenAndServe(":8080",h)
}