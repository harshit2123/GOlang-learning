package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println("Learn about Slices")
	
	var arr = []string{"A","b","c"}
	fmt.Println("array is: ",arr)
	fmt.Printf("Type of fruit list %T\n", arr)
	arr = append(arr, "d","e")

	fmt.Println("new Array is ", arr)
	//arr = append(arr[1:3])
	// Kinda  => [x: y)
	fmt.Println("new Sliced Array is ", arr)

	hiscores := make([]int,4)

	hiscores[0] =1
	hiscores[1]=2
	hiscores[2]=3
	hiscores[3] =5

	hiscores= append(hiscores, 34,45,32)
	fmt.Println(hiscores)
	sort.Ints(hiscores)
	fmt.Println(hiscores)



	// how to remove a value from slices based on the index

	var courses = []string{"A","B","C","D"}
	fmt.Println(courses)

	var index =2;
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)




}