package main

import "fmt"

func main(){

	var age int
	print("Enter Your Age : ")
	_ , err := fmt.Scanf("%d",&age)

	if err!=nil {             // no bracket required in condtional expression
		println(err)
	}
	
    if age>=18 {
		println("You can Drive !")
	} else if age < 0 {                       
        println("You've got to be kidding me ")
	} else{
		println("You can't Drive !")
	}
 
	/*
	  you have to start else or else if block exactly next to the previous block

	  if condition {

	  } else {     " correct way  "

	  }


     if condition {

	 }
	 else{          " wrong way   "

	 }

	*/
	println(" ------------ using switch ----------")

	switch  {            
	case age>=18:
		println("You can Drive !")
		break
	case age<18 && age>=0:
		println("You can't Drive !")
		break
	default:
		println("You've got to be kidding me ")

	}


	/*

	switch expression {

	}
	 expression should be same type as case expressions 
	  if expression is not present in switch then by default it would be boolean

	*/



}