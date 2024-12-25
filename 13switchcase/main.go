package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	fmt.Println("Learning Switch case in golang")
	//seed the random number generator
	rand.Seed(time.Now().UnixNano())

	//generate a random numbre between 0 and 6
	diceNuumber:= rand.Intn(6)+1 // added 1 coz Intn(6) was non inclusive
	fmt.Println("Value of dice is", diceNuumber)

	switch diceNuumber{
	case 1:
		fmt.Println("Dice value is 1 , you can open")
	case 2:
		fmt.Println("You can move to 2 spot")
	case 3:
		fmt.Println("You can move to 3 spot")
	case 4:
		fmt.Println("You can move to 4 spot")
	case 5:
		fmt.Println("You can move to 5 spot")
	case 6:
		fmt.Println("You can move to 6 spot and rool the dice again")
	default:
		fmt.Println("What are u doin")
	}


}   