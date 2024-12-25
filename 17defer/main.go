package main

import "fmt"

// Defer is execeuted in LIFO fashion
func main(){
	fmt.Println("Learning about defer in golang ")
	defer fmt.Println("YO")
	myDefer()
}

func myDefer(){
	for i:=0; i<5;i++{
		defer fmt.Println(i)
	}
}