package main

func main(){

	// there is only for loop in Go 
   
	 for i := 0; i<10;i++{
		 println(i)
	 }
	 
	 for{
		 println("Infinite loop")
		 break
	 }
	
	 println("Odd Numbers ")
	 for i:= 0; i<10;i++{
		  if i%2==1{
			  continue
		  }
		  println(i)
	 }

}