package main

import "fmt"

func main(){
	fmt.Println("Welcome to loops in go lang")
	days :=[]string{"sun","mon","tue","wed","thur","fri","sat"}
	fmt.Println(days) 


	// for d:=0; d<len(days); d++{
	// 	fmt.Println(days[d])
	// }
	for i:=range days{
		fmt.Println(days[i])
	}
     

	extra:=1
	for extra<10{
		if extra ==2{
			goto lco
		}
		if(extra==5){
			extra++
			continue
		}
		fmt.Println("Value is: ",extra)
		extra++
	}

	lco:
	    fmt.Println("Jumping")
}