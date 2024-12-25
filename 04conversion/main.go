package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("Welcome to our App")
	fmt.Println("Please give rate to our app between 1 to 5")


	reader :=bufio.NewReader(os.Stdin)
	input,_:= reader.ReadString('\n')
	fmt.Println("Thanx for rating, ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

    if err != nil{
		fmt.Println(err)
	}else{
    fmt.Println("Added 1 to your rating: ", numRating+1)
	}

} 