/*

 []type


*/


package main

import "fmt"

func main(){
   
	var slice1 = []int{}  // empty slice
	slice2 := []string{"a","b","c"}
	slice3 := make([]int,3) // slice have three elements initially
	
	slice3[0] = 3
	slice3[1] = 2
	slice3[2] = 1

	slice1 = append(slice1,1)
	slice2 = append(slice2,"d")
	slice3 = append(slice3,-1)
	slice4 := make([]int,len(slice3))

	copy(slice4,slice3)  // copy elements of slice3 into slice4 copy(destination,source)
	
	slice5 := slice2[1:3]  // [start:end-1]  
	

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
    fmt.Println(slice5)


}