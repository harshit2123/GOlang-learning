package main

import "fmt"

func main(){
	fmt.Println("Everything about pointers in go")
	// var ptr *int 

	myNUmber := 25
	var ptr = &myNUmber //thi sis basically the reference, *ptr will give you the value

	fmt.Println("value is ", ptr)
	fmt.Println("Value is ", *ptr)

	*ptr *=2 // the operations are performed on those values not on the copis
	fmt.Println("New Value is ", myNUmber)


}