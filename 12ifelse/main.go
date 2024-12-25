package main

import "fmt"

func main(){
	fmt.Println("Learning if else")

	count:=11
	var result string

	if count<10{
		result = "Regular User"
	}else{
		result = "something else"
	}
	fmt.Println(result)


	if 9%2==0 {
        fmt.Println("Even")
	}else{
		fmt.Println("Odd")
	}
}