package main

import "fmt"

func main(){
 fmt.Println("Learning about struct now")
 // no inheritance in golang

 harshit:= User{"Harshit","harshit@go.dev", true, 22}
 fmt.Println(harshit)
//  fmt.Println("hitesh's details are: %+v\n", harshit)
//  fmt.Println("Name is %v and Email is %v.\n", harshit.Name, harshit.Email)

 harshit.GetStatus()
 harshit.NewMail()

}

type User struct{
	Name string
	Email string
	Status bool
	Age int 

}

func (u User) GetStatus(){
	fmt.Println("Is user active", u.Status)

}

func(u User) NewMail(){
	u.Email = "test@go"
	fmt.Println("Email of the user is: ",u.Email)
}