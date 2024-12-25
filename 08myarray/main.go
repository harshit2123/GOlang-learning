package main

import "fmt"

func main(){
	fmt.Println("Learning about arrays")
	var arr[4]string

	arr[0] ="A"
	arr[1] = "B"
	arr[3] = "P"
	fmt.Println("Array list", arr)
	fmt.Println("Array Size:", len(arr))

	var vegies = [3]string{"pot","beans","mushroom"}
	fmt.Println("vegies list", len(vegies))
}