package main

import "fmt"

func main(){
	fmt.Println("Learning about Functions")
	greeter()


	result:= adder(3,5)
	fmt.Println("Result is", result)
	proans := proadder(2,3,4,56,7,3)
	fmt.Println("Pro result", proans)

}
func greeter(){
	fmt.Println("namastey")
}

func adder(valone int, valtwo int) int{
   return valone+valtwo
}

//this is like i dont nkow , how many values are gonna come, 
// here values is like a slice
func proadder(values...int) int {
	total:=0
	for _,val:=range values{
		total+=val
	}
	return total
}