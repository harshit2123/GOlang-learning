package main

import "fmt"

func main(){
 fmt.Println("Learning about struct now")
 // no inheritance in golang

 harshit:= User{"Harshit","harshit@go.dev", true, 22}
 fmt.Println(harshit)
//  fmt.Println("hitesh's details are: %+v\n", harshit)
//  fmt.Println("Name is %v and Email is %v.", harshit.Name, harshit.Email)

}

type User struct{
	Name string
	Email string
	Status bool
	Age int 

}