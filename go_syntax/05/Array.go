package main

import "fmt"

func main(){
	var Arr [5]int
	fmt.Println(Arr)   //  default value of all elements is zero
	
	for i:=0 ; i<5;i++{
		Arr[i] = i;
	}

	fmt.Println(Arr)
	Arr2 := [5]int{1,2,3,4,5}
	fmt.Println(Arr2)
	Arr2D := [2][5]int{{1,2,3,4,5},{6,7,8,9,10}}
	fmt.Println(Arr2D)
}