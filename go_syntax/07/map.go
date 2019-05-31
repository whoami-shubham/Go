/*
  
   map[key_type]value_type

*/

package main

import "fmt"

func main(){
   
	map1 := map[string]int{"one":1,"two":2,"three":3}
	map2 := make(map[int]string)
	map2[1] = "one"
	map2[2] = "two"
	map2[3] = "three"
	map1["four"] = 4

	fmt.Println(map1)
	fmt.Println(map2)
   
	delete(map1,"four")     // delete(map,key)

	fmt.Println(map1)

	value,key_is_present := map1["four"]   
	fmt.Println(value,key_is_present)
	value,key_is_present  = map1["three"]
	fmt.Println(value,key_is_present)

    



}