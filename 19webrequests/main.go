package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://quest.squadcast.tech/api/2021UCH1682/weather"


func main(){
	fmt.Println("Learning about web requests in Golang")
	response, err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close() // always close the connection

	databytes,err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
 // As databytes are in the numbers 
	content := string(databytes)
	fmt.Println(content)

}

func checkNilErr(err error){
	if err!=nil{
		panic(err)
	}
}